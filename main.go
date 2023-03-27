package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Algorithm interface {
	Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int
}

func main() {

	// matriz1 := readMatrix("matriz3.txt")
	// matriz2 := matriz1
	// matriz3 := make([][]int, len(matriz1))

	// for i := range matriz3 {
	// 	matriz3[i] = make([]int, len(matriz2[0]))
	// }

	// imprimirMatriz(matriz1)

	// var algorithms []Algorithm
	// a1 := A1_NaivStandard{}

	// // algorithms = append(algorithms, a1)

	// matriz3 = a1.NaivStandard(matriz1, matriz2, matriz3)

	// imprimirMatriz(matriz3)
	// algorithms = append(algorithms, A2_NaivOnArray{})
	// algorithms = append(algorithms, A3_NaivKahan{})
	// algorithms = append(algorithms, A4_NaivLoopUnrollingTwo{})
	// algorithms = append(algorithms, A5_NaivLoopUnrollingThree{})
	// algorithms = append(algorithms, A6_NaivLoopUnrollingFour{})
	// algorithms = append(algorithms, A7_WinogradOriginal{})
	// algorithms = append(algorithms, A8_WinogradScaled{})
	// algorithms = append(algorithms, A9_StrassenNaiv{}) //está haciendo mal la multiplicación
	// algorithms = append(algorithms, A10_StrassenWinograd{})     //está haciendo mal la multiplicación
	// algorithms = append(algorithms, A11_III_3SequentialBlock{}) //hay que revisarlo
	// algorithms = append(algorithms, A12_III_4ParallelBlock{})   //hay que revisarlo

	// for _, algorithm := range algorithms {
	// 	fmt.Println("\nAlgoritmo: ")
	// 	matrixSize := len(matriz1)
	// 	//nombre del algoritmo pero sin el "main."
	// 	algorithmName := strings.Split(fmt.Sprintf("%T", algorithm), ".")[1]
	// 	fmt.Println(algorithmName + " " + strconv.Itoa(matrixSize) + "x" + strconv.Itoa(matrixSize))
	// 	//matriz3 = algorithm.Run(matriz1, matriz2, matriz3)
	// 	//imprimirMatriz(matriz3)
	// }

	url := "https://script.google.com/macros/s/AKfycbxr7hkHFBYpIedsbSWq91jGiOnMbk1iGTYaCbG5VsHAcWB8IphUbL1_lJjysd4Zu57ZIg/exec?"

	idAlgo := "4"

	columna := "4"

	tiempo := "3"

	tamanoMatriz := "2"

	url = url + "idAlgo=" + idAlgo + "&" + "tamMatriz=" + columna + "&" + "tiempo=" + tiempo + "&" + "tamanoMatriz=" + tamanoMatriz

	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al enviar solicitud:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Código de respuesta:", resp.Status)

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
