package main

type A13_IV_3SequentialBlockstruct struct {
}

func multiply(A [][]int, B [][]int) [][]int {
	size := len(A)
	C := make([][]int, size)
	for i := range C {
		C[i] = make([]int, size)
	}
	bsize := 32
	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				for i := i1; i < min(i1+bsize, size); i++ {
					for j := j1; j < min(j1+bsize, size); j++ {
						for k := k1; k < min(k1+bsize, size); k++ {
							C[i][k] += A[i][j] * B[j][k]
						}
					}
				}
			}
		}
	}
	return C
}
