package main

type A1_NaivStandard struct {
}

func (a A1_NaivStandard) NaivStandard(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {

	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz2[0]); j++ {
			for k := 0; k < len(matriz1[0]); k++ {
				matriz3[i][j] += matriz1[i][k] * matriz2[k][j]
			}
		}
	}
	return matriz3
}

func (a A1_NaivStandard) Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {
	return a.NaivStandard(matriz1, matriz2, matriz3)
}
