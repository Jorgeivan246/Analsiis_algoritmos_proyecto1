package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type algoritmo interface {
	Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int
}

func main() {

	var algoritmos []algoritmo
	algoritmos = append(algoritmos, A1_NaivStandard{})
	algoritmos = append(algoritmos, A2_NaivOnArray{})
	algoritmos = append(algoritmos, A3_NaivKahan{})
	algoritmos = append(algoritmos, A4_NaivLoopUnrollingTwo{})
	algoritmos = append(algoritmos, A5_NaivLoopUnrollingThree{})
	algoritmos = append(algoritmos, A6_NaivLoopUnrollingFour{})
	algoritmos = append(algoritmos, A7_WinogradOriginal{})
	algoritmos = append(algoritmos, A8_WinogradScaled{})
	// algoritmos = append(algoritmos, A9_StrassenNaiv{}) //está haciendo mal la multiplicación
	// algoritmos = append(algoritmos, A10_StrassenWinograd{})     //está haciendo mal la multiplicación
	// algoritmos = append(algoritmos, A11_III_3SequentialBlock{}) //hay que revisarlo
	// algoritmos = append(algoritmos, A12_III_4ParallelBlock{})   //hay que revisarlo

	enviarDatosAlServidor(algoritmos)

}

func enviarDatosAlServidor(algoritmos []algoritmo) {

	url := "https://script.google.com/macros/s/AKfycbxODvkzedb9yE9Unwtj6sf6x0AS27mB2Mt3UHMKLABoIsCX3KxW-v7pa0-F_3sSF5UZmw/exec"

	var tiempo float64 = 0.000000

	var idAlgoritmo = 4

	var matriz1 [][]int

	var matriz2 [][]int

	var matriz3 [][]int

	var cantidadDatosPrueba = 4

	var tamanoMatrizAleer = 1

	var tamanoMatriz2 = 0

	var columna = 4

	for _, algoritmo := range algoritmos {
		fmt.Println("\nAlgoritmo: ")
		matrixSize := len(matriz1)
		//nombre del algoritmo pero sin el "main."
		algoritmoName := strings.Split(fmt.Sprintf("%T", algoritmo), ".")[1]
		fmt.Println(algoritmoName + " " + strconv.Itoa(matrixSize) + "x" + strconv.Itoa(matrixSize))
		tamanoMatrizAleer = 1
		for i := 0; i < cantidadDatosPrueba; i++ {

			matriz3, matriz2, matriz1, tamanoMatriz2 = inicializarMatrizTamanoIgual(tamanoMatrizAleer, matriz1, matriz2, matriz3)

			start := time.Now()
			matriz3 = algoritmo.Run(matriz1, matriz2, matriz3)

			elapsed := time.Since(start)

			elapsedSeconds := float64(elapsed) / float64(time.Second)

			tiempo = elapsedSeconds

			tiempo2 := strconv.FormatFloat(tiempo, 'f', 6, 64)

			fmt.Println(tiempo2)

			fmt.Println(idAlgoritmo, columna)

			url = url + "?" + "idAlgo=" + strconv.Itoa(idAlgoritmo) + "&" + "columna=" + strconv.Itoa(columna) + "&" + "tamanoMatriz=" + strconv.Itoa(tamanoMatriz2) + "&" + "tiempo=" + tiempo2

			fmt.Println(url)

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error al enviar solicitud:", err)
				return
			}

			defer resp.Body.Close()

			fmt.Println("Código de respuesta:", resp.Status)
			tamanoMatrizAleer = tamanoMatrizAleer + 1

			columna = columna + 2

			url = "https://script.google.com/macros/s/AKfycbxODvkzedb9yE9Unwtj6sf6x0AS27mB2Mt3UHMKLABoIsCX3KxW-v7pa0-F_3sSF5UZmw/exec"

		}

		idAlgoritmo = idAlgoritmo + 1
		columna = 4
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
