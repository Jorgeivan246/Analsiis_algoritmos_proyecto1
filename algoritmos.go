package main

import (
	"fmt"
	"math/rand"
)

func main() {

	matriz1 := crearMatriz(4, 3)
	matriz2 := crearMatriz(4, 5)

	matriz3 := make([][]int, len(matriz1))

	for i := range matriz3 {
		matriz3[i] = make([]int, len(matriz2[0]))
	}

	matriz3 = metodoNaivStandard(matriz1, matriz2, matriz3)

	println("Se imprime la matriz 3 con el metodo metodoNaivStandard")
	imprimirMatriz(matriz3)

	matriz3 = naivOnArray(matriz1, matriz2, matriz3)

	fmt.Println("\nSe imprime la matriz 3 con el metodo naivOnArray")
	imprimirMatriz(matriz3)
}

func metodoNaivStandard(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {

	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz2[0]); j++ {
			for k := 0; k < len(matriz1[0]); k++ {
				matriz3[i][j] += matriz1[i][k] * matriz2[k][j]
			}
		}
	}

	return matriz3

}

func naivOnArray(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz2[0]); j++ {
			result[i][j] = 0
			for k := 0; k < len(matriz1[0]); k++ {
				result[i][j] += matriz1[i][k] * matriz2[k][j]
			}
		}
	}
	return result
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

// Metodo para llenar una matriz con numeros aleatorios
func llenarMatriz(matriz [][]int) [][]int {
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[0]); j++ {
			matriz[i][j] = rand.Intn(10)
		}
	}
	return matriz
}
