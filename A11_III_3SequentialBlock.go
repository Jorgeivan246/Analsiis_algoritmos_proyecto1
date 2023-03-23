package main

type A11_III_3SequentialBlock struct {
}

/**
Este método aplica el metodo SequentialBlock numero 11
"bsize" es el tamaño de bloque utilizado para optimizar el rendimiento de la multiplicación de matrices.
**/

func (s A11_III_3SequentialBlock) sequentialBlock(A [][]int, B [][]int, size int, bsize int) [][]int {
	// Crear la matriz "C" de tamaño "size".
	C := make([][]int, size)
	for i := range C {
		// Inicializar cada fila de la matriz "C" con un slice de tamaño "size".
		C[i] = make([]int, size)
	}

	// Iterar sobre los bloques de tamaño "bsize" en las matrices "A", "B" y "C".
	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				// Iterar sobre cada entrada en los bloques actuales.
				for i := i1; i < i1+bsize && i < size; i++ {
					for j := j1; j < j1+bsize && j < size; j++ {
						for k := k1; k < k1+bsize && k < size; k++ {
							// Calcular la entrada correspondiente en la matriz resultante "C".
							C[i][j] += A[i][k] * B[k][j]
						}
					}
				}
			}
		}
	}

	// Devolver la matriz resultante "C".
	return C
}

func (A11_III_3SequentialBlock) Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {
	return A11_III_3SequentialBlock{}.sequentialBlock(matriz1, matriz2, len(matriz1), 3)
}