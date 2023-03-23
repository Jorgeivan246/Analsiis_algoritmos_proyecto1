package main

type A12_III_4ParallelBlock struct {
}

/*
*

# Esta funcion aplica el metodo Parallel Block numero 12

*
*/
func (p A12_III_4ParallelBlock) matrixMultiplication(A, B, C [][]float64) [][]float64 {
	size := len(A)
	bsize := 32

	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				for i := i1; i < min(i1+bsize, size); i++ {
					for j := j1; j < min(j1+bsize, size); j++ {
						for k := k1; k < min(k1+bsize, size); k++ {
							A[k][i] += B[k][j] * C[j][i]
						}
					}
				}
			}
		}
	}

	return A
}

/*
*

# Esta funcion es auxiliar del metodo numero 11 Parallel Block

*
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
