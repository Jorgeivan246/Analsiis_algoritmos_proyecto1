package main

type IV_3_SequentialBlock struct {
}

func (v IV_3_SequentialBlock) SequentialBlockMetodo(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
	size := len(matriz1)
	bsize := 16 // valor asumido para el tamaño de bloque

	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				for i := i1; i < min(i1+bsize, size); i++ {
					for j := j1; j < min(j1+bsize, size); j++ {
						for k := k1; k < min(k1+bsize, size); k++ {
							result[k][i] += matriz1[k][j] * matriz2[j][i]
						}
					}
				}
			}
		}
	}

	return result
}
