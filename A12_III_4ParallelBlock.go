package main

import "sync"

type A12_III_4ParallelBlock struct {
}

/*
*

# Esta funcion aplica el metodo Parallel Block numero 12

*
*/
func (p A12_III_4ParallelBlock) A12_III_4ParallelBlock(A [][]int, B [][]int, C [][]int) [][]int {
	size := len(A)
	bsize := 2
	var wg sync.WaitGroup
	wg.Add(size / bsize * size / bsize * size / bsize)
	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				go func(i1, j1, k1 int) {
					defer wg.Done()
					for i := i1; i < i1+bsize && i < size; i++ {
						for j := j1; j < j1+bsize && j < size; j++ {
							for k := k1; k < k1+bsize && k < size; k++ {
								A[i][k] += B[i][j] * C[j][k]
							}
						}
					}
				}(i1, j1, k1)
			}
		}
	}
	wg.Wait()

	return C
}

func (A12_III_4ParallelBlock) Run(A [][]int, B [][]int, C [][]int) [][]int {
	return A12_III_4ParallelBlock{}.A12_III_4ParallelBlock(A, B, C)
}
