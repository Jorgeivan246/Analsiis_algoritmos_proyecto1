package main

import (
	"fmt"
	"strconv"
)

type A11_III_3SequentialBlock struct {
}

/**
Este método aplica el metodo SequentialBlock numero 11
"bsize" es el tamaño de bloque utilizado para optimizar el rendimiento de la multiplicación de matrices.
**/

func (s A11_III_3SequentialBlock) sequentialBlock(A [][]int, B [][]int, C [][]int) [][]int {
	// Crear la matriz "C" de tamaño "size".

	var bsize = obtenerLongitudBsize(A)

	for i1 := 0; i1 < len(A); i1 += bsize {
		for j1 := 0; j1 < len(B); j1 += bsize {
			for k1 := 0; k1 < len(A); k1 += bsize {
				for i := i1; i < i1+bsize && i < len(A); i++ {
					for j := j1; j < j1+bsize && j < len(A); j++ {
						for k := k1; k < k1+bsize && k < len(A); k++ {
							A[i][j] += B[i][k] * C[k][j]
						}
					}
				}
			}
		}
	}

	// Devolver la matriz resultante "C".
	return C
}

func obtenerLongitudBsize(matriz1 [][]int) int {

	var mapa_bsize map[string]int

	mapa_bsize = make(map[string]int)

	var nMatriz = len(matriz1)

	var nMatrizAux = strconv.Itoa(nMatriz)

	mapa_bsize["2"] = 1
	mapa_bsize["4"] = 2
	mapa_bsize["8"] = 3
	mapa_bsize["16"] = 4
	mapa_bsize["32"] = 5
	mapa_bsize["64"] = 6
	mapa_bsize["128"] = 6
	mapa_bsize["256"] = 6
	mapa_bsize["512"] = 6
	mapa_bsize["1024"] = 6
	mapa_bsize["2048"] = 6
	mapa_bsize["4096"] = 6

	valor := mapa_bsize[nMatrizAux]

	fmt.Print(valor)

	return valor
}

func (A11_III_3SequentialBlock) Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {
	return A11_III_3SequentialBlock{}.sequentialBlock(matriz1, matriz2, matriz3)
}
