package main

import "fmt"

type A15_V_3SequentialBlock struct {
}

func (s A15_V_3SequentialBlock) V_3SequentialBlock(A [][]int, B [][]int, C [][]int) [][]int {

	var bsize = obtenerLongitudBsize(A)

	var size = len(A)

	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				for i := i1; i < i1+bsize && i < size; i++ {
					for j := j1; j < j1+bsize && j < size; j++ {
						for k := k1; k < k1+bsize && k < size; k++ {
							fmt.Println("i: ", i, "j: ", j, "k: ", k)
							C[k][i] += A[k][j] * B[j][i]
						}
					}
				}
			}
		}
	}

	return C
}

func (A15_V_3SequentialBlock) Run(A [][]int, B [][]int, C [][]int) [][]int {
	return A15_V_3SequentialBlock{}.V_3SequentialBlock(A, B, C)
}
