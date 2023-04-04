package main

type A7_WinogradOriginal struct {
}

func (a A7_WinogradOriginal) winogradOriginal(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
	n := len(matriz2[0])
	p := len(matriz2)
	m := len(matriz1)
	var aux int
	upsilon := p % 2
	gamma := p - upsilon
	y := make([]int, m)
	z := make([]int, n)

	for i := 0; i < m; i++ {
		aux = 0
		for j := 0; j < gamma; j += 2 {
			aux += matriz1[i][j] * matriz1[i][j+1]
		}
		y[i] = aux
	}

	for i := 0; i < n; i++ {
		aux = 0
		for j := 0; j < gamma; j += 2 {
			aux += matriz2[j][i] * matriz2[j+1][i]
		}
		z[i] = aux
	}

	if upsilon == 1 {
		PP := p - 1
		for i := 0; i < m; i++ {
			for k := 0; k < n; k++ {
				aux = 0
				for j := 0; j < gamma; j += 2 {
					aux += (matriz1[i][j] + matriz2[j+1][k]) * (matriz1[i][j+1] + matriz2[j][k])
				}
				result[i][k] = aux - y[i] - z[k] + matriz1[i][PP]*matriz2[PP][k]
			}
		}
	} else {
		for i := 0; i < m; i++ {
			for k := 0; k < n; k++ {
				aux = 0
				for j := 0; j < gamma; j += 2 {
					aux += (matriz1[i][j] + matriz2[j+1][k]) * (matriz1[i][j+1] + matriz2[j][k])
				}
				result[i][k] = aux - y[i] - z[k]
			}
		}
	}
	return result
}

func (a A7_WinogradOriginal) Run(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {
	return a.winogradOriginal(matriz1, matriz2, matriz3)
}