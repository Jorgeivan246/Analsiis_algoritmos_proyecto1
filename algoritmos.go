package main

import (
	"fmt"
)

func main() {

	var matriz1 [3][3]int

	var matriz2 [3][3]int

	matriz3 := make([][]int, len(matriz1))

	var a int = 1

	for i := range matriz3 {
		matriz3[i] = make([]int, len(matriz2[0]))
	}

	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz2); j++ {
			matriz1[i][j] = a
		}

	}

	for i := 0; i < len(matriz2); i++ {
		for j := 0; j < len(matriz2); j++ {
			matriz2[i][j] = a
		}
	}

	matriz3 = metodoNaivStandard(matriz1, matriz2, matriz3)

	for i := 0; i < len(matriz3); i++ {
		fmt.Println(matriz3[i])
	}
}

func metodoNaivStandard(matriz1 [3][3]int, matriz2 [3][3]int, matriz3 [][]int) [][]int {

	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz2[0]); j++ {
			for k := 0; k < len(matriz3); k++ {
				matriz3[i][j] += matriz1[i][k] * matriz2[k][j]
			}
		}
	}

	return matriz3

}
