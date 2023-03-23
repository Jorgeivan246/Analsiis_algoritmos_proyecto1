package main

import "math"

type A8_WinogradScaled struct {
}

func (w A8_WinogradScaled) winogradScaled(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
	n := len(matriz1)
	p := len(matriz1[0])
	m := len(matriz2[0])
	// Crear copias escaladas de A y B
	copyA := make([][]int, n)
	for i := 0; i < n; i++ {
		copyA[i] = make([]int, p)
	}
	copyB := make([][]int, p)
	for i := 0; i < p; i++ {
		copyB[i] = make([]int, m)
	}
	// Factores de escala
	a := w.normInf(matriz1, n, p)
	b := w.normInf(matriz2, p, m)
	lambda := int(math.Floor(0.5 + math.Log(float64(b)/float64(a))/math.Log(4)))
	// Escalado
	w.multiplyWithScalar(matriz1, copyA, n, p, int(math.Pow(2, float64(lambda))))
	w.multiplyWithScalar(matriz2, copyB, p, m, int(math.Pow(2, float64(-lambda))))
	// Usando Winograd con matrices escaladas
	aux := A7_WinogradOriginal{}
	result = aux.winogradOriginal(copyA, copyB, result)
	return result
}

/*
Esta función calcula la norma infinito de una matriz A de tamaño N x P.
La norma infinito se define como el máximo valor absoluto de la suma de cada fila de la matriz.
La función devuelve un valor entero que representa la norma infinito de la matriz.
*/
func (a A8_WinogradScaled) normInf(A [][]int, N int, P int) int {
	max := 0
	for i := 0; i < N; i++ {
		sum := 0
		for j := 0; j < P; j++ {
			sum += int(math.Abs(float64(A[i][j])))
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

/*
Esta función multiplica cada elemento de una matriz A de tamaño N x P por un escalar entero dado
y almacena el resultado en una matriz B del mismo tamaño. La función no devuelve ningún valor.
*/
func (a A8_WinogradScaled)  multiplyWithScalar(A [][]int, B [][]int, N int, P int, scalar int) {
	for i := 0; i < N; i++ {
		for j := 0; j < P; j++ {
			B[i][j] = A[i][j] * scalar
		}
	}
}

func (a A8_WinogradScaled) Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {
	return a.winogradScaled(matriz1, matriz2, matriz3)
}