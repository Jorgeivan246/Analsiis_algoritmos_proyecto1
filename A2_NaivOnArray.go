package main

type A2_NaivOnArray struct {
}

func (a A2_NaivOnArray) naiveOnArray(matriz1 [][]int, matriz2 [][]int) []int {
	result := make([]int, len(matriz1)*len(matriz2[0]))
	n := len(matriz1)
	m := len(matriz2[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result[i*m+j] = 0
			for k := 0; k < len(matriz1[0]); k++ {
				result[i*m+j] += matriz1[i][k] * matriz2[k][j]
			}
		}
	}
	return result
}

func sliceToMatrix(slice []int, n int, m int) [][]int {
	matriz := make([][]int, n)
	for i := 0; i < n; i++ {
		matriz[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matriz[i][j] = slice[i*m+j]
		}
	}
	return matriz
}

func (a A2_NaivOnArray) Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {

	return sliceToMatrix(a.naiveOnArray(matriz1, matriz2), len(matriz1), len(matriz2[0]))
}