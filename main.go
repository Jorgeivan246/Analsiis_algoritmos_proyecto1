package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type algoritmo interface {
	Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int
}

var url = "https://script.google.com/macros/s/AKfycbyxcN3gH-UONnj56hAUJSpWUTSuov2dYjyl-xN5pO9_qfdPJ0DL5QRQ1SU2ZfpuSID-BA/exec"

func main() {

	var algoritmos []algoritmo
	algoritmos = append(algoritmos, A1_NaivStandard{})
	// algoritmos = append(algoritmos, A2_NaivOnArray{})
	// algoritmos = append(algoritmos, A3_NaivKahan{})
	// algoritmos = append(algoritmos, A4_NaivLoopUnrollingTwo{})
	// algoritmos = append(algoritmos, A5_NaivLoopUnrollingThree{})
	// algoritmos = append(algoritmos, A6_NaivLoopUnrollingFour{})
	// algoritmos = append(algoritmos, A7_WinogradOriginal{})
	// algoritmos = append(algoritmos, A8_WinogradScaled{})
	// algoritmos = append(algoritmos, A9_StrassenNaiv{})
	// algoritmos = append(algoritmos, A10_StrassenWinograd{})
	algoritmos = append(algoritmos, A11_III_3SequentialBlock{})
	algoritmos = append(algoritmos, A12_III_4ParallelBlock{})
	algoritmos = append(algoritmos, A13_IV_3SequentialBlockstruct{})
	algoritmos = append(algoritmos, A14_IV_4ParallelBlock{})
	algoritmos = append(algoritmos, A15_V_3SequentialBlock{})
	algoritmos = append(algoritmos, A16_V_4ParallelBlock{})
	probarALgoritmo(algoritmos)

}

func obtenerDatosHardware() (string, string, string) {

	var info2 syscall.Sysinfo_t

	//Obtener datos en windows

	// info2, err := mem.VirtualMemory()
	// if err != nil {
	// 	panic(err)
	// }

	// err := syscall.Sysinfo(&info2)
	// if err != nil {
	// 	panic(err)
	// }

	// var memoria = info2.Total / 1024 / 1024 / 1024

	var memoria = info2.Totalram / 1024 / 1024 / 1024

	var cantidadHIlos = runtime.NumCPU()

	var nHilos = strconv.Itoa(cantidadHIlos)

	var modeloProcesador = ""

	var memoriaRam string = strconv.FormatUint(memoria, 12) + "GB"

	info, err := cpu.Info()
	if err != nil {
		panic(err)
	}

	if len(info) > 0 {
		modeloProcesador = info[0].ModelName

		modeloProcesador = regexp.MustCompile(`\s+`).ReplaceAllString(modeloProcesador, "")
		fmt.Println(modeloProcesador)
	} else {
		modeloProcesador = "No se pudo obtener información del procesador"
	}

	return modeloProcesador, nHilos, memoriaRam
}

func enviarDatosAlServidor(algoritmos []algoritmo) {

	var tiempo float64 = 0.00000000

	var urlAux string

	var idAlgoritmo = 3

	var columna = 4

	var matriz1 [][]int

	var matriz2 [][]int

	var matriz3 [][]int

	var cantidadCasosPrueba = 2

	var tamanoMatrizAleer = 1

	var tamanoMatriz2 = 0

	var modeloProcesador, nHilos, memoriaRam = obtenerDatosHardware()

	url_hardware := url + "?" + "idAlgo=" + strconv.Itoa(idAlgoritmo) + "&" + "columna=" + strconv.Itoa(columna) + "&" + "nombreProcesador=" + modeloProcesador + "&" + "nHilos=" + nHilos + "&" + "memoriaRam=" + memoriaRam

	fmt.Println(url_hardware)
	resp, err := http.Get(url_hardware)
	if err != nil {
		fmt.Println("Error al enviar solicitud:", err)
		return
	}

	defer resp.Body.Close()

	idAlgoritmo = idAlgoritmo + 8

	columna = columna + 1

	for _, algoritmo := range algoritmos {
		fmt.Println("\nAlgoritmo: ")
		matrixSize := len(matriz1)
		//nombre del algoritmo pero sin el "main."
		algoritmoName := strings.Split(fmt.Sprintf("%T", algoritmo), ".")[1]
		fmt.Println(algoritmoName + " " + strconv.Itoa(matrixSize) + "x" + strconv.Itoa(matrixSize))
		tamanoMatrizAleer = 1
		for i := 0; i < cantidadCasosPrueba; i++ {

			matriz3, matriz2, matriz1, tamanoMatriz2 = inicializarMatrizTamanoIgual(tamanoMatrizAleer, matriz1, matriz2, matriz3)

			start := time.Now()
			matriz3 = algoritmo.Run(matriz1, matriz2, matriz3)

			elapsed := time.Since(start)

			elapsedSeconds := float64(elapsed) / float64(time.Second)

			tiempo = elapsedSeconds

			tiempo2 := strconv.FormatFloat(tiempo, 'f', 8, 64)

			fmt.Println(tiempo2)

			fmt.Println(idAlgoritmo, columna)

			urlAux = url + "?" + "idAlgo=" + strconv.Itoa(idAlgoritmo) + "&" + "columna=" + strconv.Itoa(columna) + "&" + "tamanoMatriz=" + strconv.Itoa(tamanoMatriz2) + "&" + "tiempo=" + tiempo2

			fmt.Println(urlAux)

			resp, err := http.Get(urlAux)
			if err != nil {
				fmt.Println("Error al enviar solicitud:", err)
				return
			}

			defer resp.Body.Close()

			fmt.Println("Código de respuesta:", resp.Status)
			tamanoMatrizAleer = tamanoMatrizAleer + 1

			columna = columna + 2

			urlAux = url
		}

		idAlgoritmo = idAlgoritmo + 1
		columna = 5
	}

}

func probarALgoritmo(algoritmos []algoritmo) {

	var matriz1 [][]int

	var matriz2 [][]int

	var matriz3 [][]int

	var tamanoMatrizAleer = 1

	tamano := float64(math.Pow(2, float64(tamanoMatrizAleer)))

	tamanoEntero := int(tamano)

	nommbreMatriz := "matriz" + strconv.Itoa(tamanoEntero) + ".txt"

	matriz1 = readMatrix(nommbreMatriz)

	matriz2 = readMatrix(nommbreMatriz)

	matriz3 = make([][]int, len(matriz1))

	for i := range matriz3 {
		matriz3[i] = make([]int, len(matriz1[i]))

	}

	for _, algoritmo := range algoritmos {

		matriz3 = algoritmo.Run(matriz1, matriz2, matriz3)
		imprimirMatriz(matriz3)

	}

}

func inicializarMatrizTamanoIgual(tamanoMatriz int, matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) ([][]int, [][]int, [][]int, int) {
	tamano := float64(math.Pow(2, float64(tamanoMatriz)))

	tamanoEntero := int(tamano)

	fmt.Println("Esta leyendo " + strconv.Itoa(tamanoEntero))
	nommbreMatriz := "matriz" + strconv.Itoa(tamanoEntero) + ".txt"

	matriz1 = readMatrix(nommbreMatriz)

	matriz2 = make([][]int, len(matriz1))

	for i := range matriz1 {
		matriz2[i] = make([]int, len(matriz1[i]))
		copy(matriz2[i], matriz1[i])
	}

	matriz3 = make([][]int, len(matriz1))

	for i := range matriz3 {
		matriz3[i] = make([]int, len(matriz2[0]))
	}
	tamanoMatriz = tamanoMatriz + 1

	return matriz3, matriz2, matriz1, tamanoEntero
}

// Metodo para crear una matriz vacia de tamaño n x m
func crearMatriz(n int, m int) [][]int {
	matriz := make([][]int, n)
	for i := 0; i < n; i++ {
		matriz[i] = make([]int, m)
	}
	//Se llena la matriz con numeros aleatorios
	matriz = llenarMatriz(matriz)
	return matriz
}

// Metodo para imprimir una matriz
func imprimirMatriz(matriz [][]int) {
	fmt.Println("Se imprime la matriz")
	for i := 0; i < len(matriz); i++ {
		fmt.Println(matriz[i])
	}
}

// Metodo para llenar una matriz con solo el numero 3
func llenarMatriz(matriz [][]int) [][]int {
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[0]); j++ {
			//se llena la matriz con numeros aleatorios
			//matriz[i][j] = rand.Intn(10)
			matriz[i][j] = 3
		}
	}
	return matriz
}

/*
* Genera una matriz cuadrada de tamaño n x n con números aleatorios
* de 4 digitos, y los almacena en un archivo .txt
* @param n = cantidad de gilas y columnas de la matriz
* @param filename = direccion y nombre del archivo.
 */
func generateMatrix(n int, filename string) {
	dir := filepath.Dir("matrices/" + filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	file, err := os.Create("matrices/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprintf(file, "%04d ", rand.Intn(9000)+1000)
		}
		fmt.Fprintln(file)
	}
}

func readMatrix(filename string) [][]int {
	file, err := os.Open("matrices/" + filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]int
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var row []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}

	return matrix
}
