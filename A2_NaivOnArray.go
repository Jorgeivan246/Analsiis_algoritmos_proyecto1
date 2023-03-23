package main

type A2_NaivOnArray struct {
}

func (a A2_NaivOnArray) naiveOnArray(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz2[0]); j++ {
			result[i][j] = 0
			for k := 0; k < len(matriz1[0]); k++ {
				result[i][j] += matriz1[i][k] * matriz2[k][j]
			}
		}
	}
	return result
}

func (a A2_NaivOnArray) Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {
	return a.naiveOnArray(matriz1, matriz2, matriz3)
}