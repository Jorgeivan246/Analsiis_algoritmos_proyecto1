package main

import "sync"

type A14_IV_4ParallelBlock struct {
}

/*
Metodo igual que el anterior que agrega una go rutina cada vez que
Va a iterar sobre un bloque
*/

func (s A14_IV_4ParallelBlock) ParallelBlock(A [][]int, B [][]int, C [][]int) [][]int {

	size := len(A)
	bsize := obtenerLongitudBsize(A)
	var wg sync.WaitGroup

	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {

				wg.Add(1)
				go func(i1, j1, k1 int) {

					for i := i1; i < i1+bsize && i < size; i++ {
						for j := j1; j < j1+bsize && j < size; j++ {
							for k := k1; k < k1+bsize && k < size; k++ {
								C[i][k] += A[i][j] * B[j][k]
							}
						}
					}

					defer wg.Done()
				}(i1, j1, k1)
			}
		}
	}
	wg.Wait()

	return C
}

func (A14_IV_4ParallelBlock) Run(A [][]int, B [][]int, C [][]int) [][]int {
	return A14_IV_4ParallelBlock{}.ParallelBlock(A, B, C)
}
