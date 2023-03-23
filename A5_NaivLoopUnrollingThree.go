package main

type A5_NaivLoopUnrollingThree struct {
}

func (a A5_NaivLoopUnrollingThree) naiveLoopUnrollingThree(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
	var i int
	var j int
	var k int
	var aux int
	if len(matriz1[0])%3 == 0 {
		for i = 0; i < len(matriz1); i++ {
			for j = 0; j < len(matriz2[0]); j++ {
				aux = 0
				for k = 0; k < len(matriz1[0]); k += 3 {
					aux += matriz1[i][k]*matriz2[k][j] + matriz1[i][k+1]*matriz2[k+1][j] + matriz1[i][k+2]*matriz2[k+2][j]
				}
				result[i][j] = aux
			}
		}
	} else if len(matriz1[0])%3 == 1 {
		var pp int = len(matriz1[0]) - 1
		for i = 0; i < len(matriz1); i++ {
			for j = 0; j < len(matriz2[0]); j++ {
				aux = 0
				for k = 0; k < pp; k += 3 {
					aux += matriz1[i][k]*matriz2[k][j] + matriz1[i][k+1]*matriz2[k+1][j] + matriz1[i][k+2]*matriz2[k+2][j]
				}
				result[i][j] = aux + matriz1[i][pp]*matriz2[pp][j]
			}
		}
	} else {
		var pp int = len(matriz1[0]) - 2
		var ppp int = len(matriz1[0]) - 1
		for i = 0; i < len(matriz1); i++ {
			for j = 0; j < len(matriz2[0]); j++ {
				aux = 0
				for k = 0; k < pp; k += 3 {
					aux += matriz1[i][k]*matriz2[k][j] + matriz1[i][k+1]*matriz2[k+1][j] + matriz1[i][k+2]*matriz2[k+2][j]
				}
				result[i][j] = aux + matriz1[i][pp]*matriz2[pp][j] + matriz1[i][ppp]*matriz2[ppp][j]
			}
		}
	}
	return result
}

func (a A5_NaivLoopUnrollingThree) Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {
	return a.naiveLoopUnrollingThree(matriz1, matriz2, matriz3)
}