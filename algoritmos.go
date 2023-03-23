package main

import (
	"fmt"
)

type Algorithm interface {
	Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int
}

func main() {

	var algorithms []Algorithm
	algorithms = append(algorithms, A1_NaivStandard{})
	algorithms = append(algorithms, A2_NaivOnArray{})
	algorithms = append(algorithms, A3_NaivKahan{})
	algorithms = append(algorithms, A4_NaivLoopUnrollingTwo{})
	algorithms = append(algorithms, A5_NaivLoopUnrollingThree{})
	algorithms = append(algorithms, A6_NaivLoopUnrollingFour{})
	algorithms = append(algorithms, A7_WinogradOriginal{})
	algorithms = append(algorithms, A8_WinogradScaled{})
	algorithms = append(algorithms, A9_StrassenNaiv{}) //está haciendo mal la multiplicación
	algorithms = append(algorithms, A10_StrassenWinograd{}) //está haciendo mal la multiplicación
	// algorithms = append(algorithms, A11_III_3SequentialBlock{})
	// algorithms = append(algorithms, A12_III_4ParallelBlock{})

	matriz1 := crearMatriz(3, 3)
	matriz2 := crearMatriz(3, 3)
	matriz3 := make([][]int, len(matriz1))

	for i := range matriz3 {
		matriz3[i] = make([]int, len(matriz2[0]))
	}

	for _, algorithm := range algorithms {
		fmt.Println("\nAlgoritmo: ")
		matriz3 = algorithm.Run(matriz1, matriz2, matriz3)
		imprimirMatriz(matriz3)
	}
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
