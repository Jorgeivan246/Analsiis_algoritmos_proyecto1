package main

import (
	"fmt"
)

func main() {

	var matriz1 [3][3]int

	var matriz2 [3][3]int

	matriz3 := make([][]int, len(matriz1))

	for i := range matriz3 {
		matriz3[i] = make([]int, len(matriz2[0]))
	}

	for i := 0; i < len(matriz3); i++ {
		fmt.Println(matriz3[i])
	}

}

func metodoNaivStandard(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {

	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz2[0]); j++ {
			for k := 0; k < len(matriz3); k++ {
				matriz3[i][j] += matriz1[i][k] * matriz2[k][j]
			}
		}
	}

	return matriz3

}
