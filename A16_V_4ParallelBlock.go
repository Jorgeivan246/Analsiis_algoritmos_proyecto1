package main

import "sync"

type A16_V_4ParallelBlock struct {
}

func (s A16_V_4ParallelBlock) V_4ParallelBlock(A [][]int, B [][]int, C [][]int) [][]int {

	size := len(A)
	bsize := obtenerLongitudBsize(A)
	var wg sync.WaitGroup

	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				// Use goroutine to execute each block in parallel
				wg.Add(1)
				go func(i1, j1, k1 int) {
					for i := i1; i < i1+bsize && i < size; i++ {
						for j := j1; j < j1+bsize && j < size; j++ {
							for k := k1; k < k1+bsize && k < size; k++ {
								C[k][i] += A[k][j] * B[j][i]
							}
						}
					}
				}(i1, j1, k1)
			}
		}
	}
	// Wait for all goroutines to finish before returning

	return C
}

func (A16_V_4ParallelBlock) Run(A [][]int, B [][]int, C [][]int) [][]int {
	return A16_V_4ParallelBlock{}.V_4ParallelBlock(A, B, C)
}
