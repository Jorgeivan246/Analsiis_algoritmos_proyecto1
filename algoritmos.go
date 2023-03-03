package main

import (
	"fmt"
)

func main() {

	matriz1 := crearMatriz(3, 3)
	matriz2 := crearMatriz(3, 3)

	matriz3 := make([][]int, len(matriz1))

	for i := range matriz3 {
		matriz3[i] = make([]int, len(matriz2[0]))
	}

	matriz3 = metodoNaivStandard(matriz1, matriz2, matriz3)

	fmt.Println("Se imprime la matriz 3 con el metodo metodoNaivStandard")
	imprimirMatriz(matriz3)

	matriz3 = naivOnArray(matriz1, matriz2, matriz3)

	fmt.Println("\nSe imprime la matriz 3 con el metodo naivOnArray")
	imprimirMatriz(matriz3)

	matriz3 = naivKahan(matriz1, matriz2, matriz3)
	fmt.Println("\nSe imprime la matriz 3 con el metodo naivKahan")
	imprimirMatriz(matriz3)

	matriz3 = naivLoopUnrollingTwo(matriz1, matriz2, matriz3)
	fmt.Println("\nSe imprime la matriz 3 con el metodo naivLoopUnrollingTwo")
	imprimirMatriz(matriz3)

	matriz3 = naivLoopUnrollingThree(matriz1, matriz2, matriz3)
	fmt.Println("\nSe imprime la matriz 3 con el metodo NaivLoopUnrollingThree")
	imprimirMatriz(matriz3)

	matriz3 = naivLoopUnrollingFour(matriz1, matriz2, matriz3)
	fmt.Println("\nSe imprime la matriz 3 con el metodo NaivLoopUnrollingFour")
	imprimirMatriz(matriz3)

	matriz3 = winogradOriginal(matriz1, matriz2, matriz3)
	fmt.Println("\nSe imprime la matriz 3 con el metodo winogradOriginal")
	imprimirMatriz(matriz3)
}

func metodoNaivStandard(matriz1 [][]int, matriz2 [][]int, matriz3 [][]int) [][]int {

	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz2[0]); j++ {
			for k := 0; k < len(matriz1[0]); k++ {
				matriz3[i][j] += matriz1[i][k] * matriz2[k][j]
			}
		}
	}

	return matriz3

}

func naivOnArray(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
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

func naivKahan(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
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

func naivLoopUnrollingTwo(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
	var i int
	var j int
	var k int
	var aux int
	if len(matriz1[0])%2 == 0 {
		for i = 0; i < len(matriz1); i++ {
			for j = 0; j < len(matriz2[0]); j++ {
				aux = 0
				for k = 0; k < len(matriz1[0]); k += 2 {
					aux += matriz1[i][k]*matriz2[k][j] + matriz1[i][k+1]*matriz2[k+1][j]
				}
				result[i][j] = aux
			}
		}
	} else {
		var pp int = len(matriz1[0]) - 1
		for i = 0; i < len(matriz1); i++ {
			for j = 0; j < len(matriz2[0]); j++ {
				aux = 0
				for k = 0; k < pp; k += 2 {
					aux += matriz1[i][k]*matriz2[k][j] + matriz1[i][k+1]*matriz2[k+1][j]
				}
				result[i][j] = aux + matriz1[i][pp]*matriz2[pp][j]
			}
		}
	}
	return result
}

func naivLoopUnrollingThree(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
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

func naivLoopUnrollingFour(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
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

func winogradOriginal(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
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
		for i := m; i < m; i++ {
			for k := n; k < n; k++ {
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

// Metodo para crear una matriz vacia de tamaño n x m
func crearMatriz(n int, m int) [][]int {
	matriz := make([][]int, n)
	for i := 0; i < n; i++ {
		matriz[i] = make([]int, m)
	}
	//Se llena la matriz con numeros aleatorios
	matriz = llenarMatriz(matriz)
	return matriz
}

// Metodo para imprimir una matriz
func imprimirMatriz(matriz [][]int) {
	for i := 0; i < len(matriz); i++ {
		fmt.Println(matriz[i])
	}
}

// Metodo para llenar una matriz con solo el numero 3
func llenarMatriz(matriz [][]int) [][]int {
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[0]); j++ {
			//se llena la matriz con numeros aleatorios
			//matriz[i][j] = rand.Intn(10)
			matriz[i][j] = 3
		}
	}
	return matriz
}
