package main

type A15_V_3SequentialBlock struct {
}

func multiply15(A, B, C [][]int, size, bsize int) {

	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				for i := i1; i < i1+bsize && i < size; i++ {
					for j := j1; j < j1+bsize && j < size; j++ {
						for k := k1; k < k1+bsize && k < size; k++ {
							A[k][i] += B[k][j] * C[j][i]
						}
					}
				}
			}
		}
	}

}
