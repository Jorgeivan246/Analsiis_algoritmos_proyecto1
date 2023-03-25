package main

import "runtime"

type A16_V_4ParallelBlock struct {
}

func multiply16(A, B, C [][]float64, size, bsize int) {
	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				// Use goroutine to execute each block in parallel
				go func(i1, j1, k1 int) {
					for i := i1; i < i1+bsize && i < size; i++ {
						for j := j1; j < j1+bsize && j < size; j++ {
							for k := k1; k < k1+bsize && k < size; k++ {
								A[k][i] += B[k][j] * C[j][i]
							}
						}
					}
				}(i1, j1, k1)
			}
		}
	}
	// Wait for all goroutines to finish before returning
	runtime.Gosched()
}
