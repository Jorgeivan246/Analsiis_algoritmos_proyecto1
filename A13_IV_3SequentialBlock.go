package main

type A13_IV_3SequentialBlockstruct struct {
}

func (s A13_IV_3SequentialBlockstruct) IV_3SequentialBlock(A [][]int, B [][]int, C [][]int) [][]int {
	// Crear la matriz "C" de tamaño "size".

	var bsize = obtenerLongitudBsize(A)

	var size = len(A)

	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				for i := i1; i < i1+bsize && i < size; i++ {
					for j := j1; j < j1+bsize && j < size; j++ {
						for k := k1; k < k1+bsize && k < size; k++ {
							C[i][k] += A[i][j] * B[j][k]
						}
					}
				}
			}
		}
	}

	// Devolver la matriz resultante "C".
	return C
}

func (A13_IV_3SequentialBlockstruct) Run(A [][]int, B [][]int, C [][]int) [][]int {
	return A13_IV_3SequentialBlockstruct{}.IV_3SequentialBlock(A, B, C)
}
