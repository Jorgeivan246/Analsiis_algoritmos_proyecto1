package main

type A6_NaivLoopUnrollingFour struct {
}

func (a A6_NaivLoopUnrollingFour) naiveLoopUnrollingFour(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
	n := len(matriz1)
	p := len(matriz1[0])
	m := len(matriz2[0])
	var i, j, k int
	var aux int
	if p%4 == 0 {
		for i = 0; i < n; i++ {
			for j = 0; j < m; j++ {
				aux = 0
				for k = 0; k < p; k += 4 {
					aux += matriz1[i][k]*matriz2[k][j] + matriz1[i][k+1]*matriz2[k+1][j] + matriz1[i][k+2]*matriz2[k+2][j] + matriz1[i][k+3]*matriz2[k+3][j]
				}
				result[i][j] = aux
			}
		}
	} else if p%4 == 1 {
		pp := p - 1
		for i = 0; i < n; i++ {
			for j = 0; j < m; j++ {
				aux = 0
				for k = 0; k < pp; k += 4 {
					aux += matriz1[i][k]*matriz2[k][j] + matriz1[i][k+1]*matriz2[k+1][j] + matriz1[i][k+2]*matriz2[k+2][j] + matriz1[i][k+3]*matriz2[k+3][j]
				}
				result[i][j] = aux + matriz1[i][pp]*matriz2[pp][j]
			}
		}
	} else if p%4 == 2 {
		pp := p - 2
		ppp := p - 1
		for i = 0; i < n; i++ {
			for j = 0; j < m; j++ {
				aux = 0
				for k = 0; k < pp; k += 4 {
					aux += matriz1[i][k]*matriz2[k][j] + matriz1[i][k+1]*matriz2[k+1][j] + matriz1[i][k+2]*matriz2[k+2][j] + matriz1[i][k+3]*matriz2[k+3][j]
				}
				result[i][j] = aux + matriz1[i][pp]*matriz2[pp][j] + matriz1[i][ppp]*matriz2[ppp][j]
			}
		}
	} else {
		pp := p - 3
		ppp := p - 2
		pppp := p - 1
		for i = 0; i < n; i++ {
			for j = 0; j < m; j++ {
				aux = 0
				for k = 0; k < pp; k += 4 {
					aux += matriz1[i][k]*matriz2[k][j] + matriz1[i][k+1]*matriz2[k+1][j] + matriz1[i][k+2]*matriz2[k+2][j] + matriz1[i][k+3]*matriz2[k+3][j]
				}
				result[i][j] = aux + matriz1[i][pp]*matriz2[pp][j] + matriz1[i][ppp]*matriz2[ppp][j] + matriz1[i][pppp]*matriz2[pppp][j]
			}
		}
	}
	return result
}

func (a A6_NaivLoopUnrollingFour) Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {
	return a.naiveLoopUnrollingFour(matriz1, matriz2, matriz3)
}