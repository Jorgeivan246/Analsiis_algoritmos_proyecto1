package main

type A3_NaivKahan struct {
}

func (a A3_NaivKahan) naiveKahan(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
	var t int
	var sum int
	var err int
	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz2); j++ {
			sum = 0
			err = 0
			for k := 0; k < len(matriz1[0]); k++ {
				err = err + matriz1[i][k]*matriz2[k][j]
				t = sum + err
				err = (sum - t) + err
				sum = t
			}
			result[i][j] = sum
		}
	}
	return result
}

func (a A3_NaivKahan) Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {
	return a.naiveKahan(matriz1, matriz2, matriz3)
}