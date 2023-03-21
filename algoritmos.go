package main

import (
	"fmt"
	"math"
)

func main() {

	matriz1 := crearMatriz(3, 3)
	matriz2 := crearMatriz(3, 3)

	matriz3 := make([][]int, len(matriz1))

	for i := range matriz3 {
		matriz3[i] = make([]int, len(matriz2[0]))
	}

	matriz3 = metodoNaivStandard(matriz1, matriz2, matriz3)

	// fmt.Println("Se imprime la matriz 3 con el metodo metodoNaivStandard")
	// imprimirMatriz(matriz3)

	// matriz3 = naivOnArray(matriz1, matriz2, matriz3)

	// fmt.Println("\nSe imprime la matriz 3 con el metodo naivOnArray")
	// imprimirMatriz(matriz3)

	// matriz3 = naivKahan(matriz1, matriz2, matriz3)
	// fmt.Println("\nSe imprime la matriz 3 con el metodo naivKahan")
	// imprimirMatriz(matriz3)

	// matriz3 = naivLoopUnrollingTwo(matriz1, matriz2, matriz3)
	// fmt.Println("\nSe imprime la matriz 3 con el metodo naivLoopUnrollingTwo")
	// imprimirMatriz(matriz3)

	// matriz3 = naivLoopUnrollingThree(matriz1, matriz2, matriz3)
	// fmt.Println("\nSe imprime la matriz 3 con el metodo NaivLoopUnrollingThree")
	// imprimirMatriz(matriz3)

	// matriz3 = naivLoopUnrollingFour(matriz1, matriz2, matriz3)
	// fmt.Println("\nSe imprime la matriz 3 con el metodo NaivLoopUnrollingFour")
	// imprimirMatriz(matriz3)

	// matriz3 = winogradOriginal(matriz1, matriz2, matriz3)
	// fmt.Println("\nSe imprime la matriz 3 con el metodo winogradOriginal")
	// imprimirMatriz(matriz3)

	// matriz3 = winogradScaled(matriz1, matriz2, matriz3)
	// fmt.Println("\nSe imprime la matriz 3 con el metodo winogradScaled")
	// imprimirMatriz(matriz3)
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

func winogradScaled(matriz1 [][]int, matriz2 [][]int, result [][]int) [][]int {
	n := len(matriz1)
	p := len(matriz1[0])
	m := len(matriz2[0])
	// Crear copias escaladas de A y B
	copyA := make([][]int, n)
	for i := 0; i < n; i++ {
		copyA[i] = make([]int, p)
	}
	copyB := make([][]int, p)
	for i := 0; i < p; i++ {
		copyB[i] = make([]int, m)
	}
	// Factores de escala
	a := normInf(matriz1, n, p)
	b := normInf(matriz2, p, m)
	lambda := int(math.Floor(0.5 + math.Log(float64(b)/float64(a))/math.Log(4)))
	// Escalado
	multiplyWithScalar(matriz1, copyA, n, p, int(math.Pow(2, float64(lambda))))
	multiplyWithScalar(matriz2, copyB, p, m, int(math.Pow(2, float64(-lambda))))
	// Usando Winograd con matrices escaladas
	result = winogradOriginal(copyA, copyB, result)
	return result
}

/*
Esta función calcula la norma infinito de una matriz A de tamaño N x P.
La norma infinito se define como el máximo valor absoluto de la suma de cada fila de la matriz.
La función devuelve un valor entero que representa la norma infinito de la matriz.
*/
func normInf(A [][]int, N int, P int) int {
	max := 0
	for i := 0; i < N; i++ {
		sum := 0
		for j := 0; j < P; j++ {
			sum += int(math.Abs(float64(A[i][j])))
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

/*
Esta función multiplica cada elemento de una matriz A de tamaño N x P por un escalar entero dado
y almacena el resultado en una matriz B del mismo tamaño. La función no devuelve ningún valor.
*/
func multiplyWithScalar(A [][]int, B [][]int, N int, P int, scalar int) {
	for i := 0; i < N; i++ {
		for j := 0; j < P; j++ {
			B[i][j] = A[i][j] * scalar
		}
	}
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

/*
Metodo número 9
Esta funcion utiliza el metodo strassenNaiv para calcular la multiplicación
*/
func strassenNaiv(a, b [][]int) [][]int {
	n := len(a)
	c := make([][]int, n)
	for i := range c {
		c[i] = make([]int, n)
	}

	if n == 1 {
		c[0][0] = a[0][0] * b[0][0]
	} else {
		// Divide las matrices en bloques más pequeños
		m := n / 2
		a11 := make([][]int, m)
		a12 := make([][]int, m)
		a21 := make([][]int, m)
		a22 := make([][]int, m)
		b11 := make([][]int, m)
		b12 := make([][]int, m)
		b21 := make([][]int, m)
		b22 := make([][]int, m)
		for i := 0; i < m; i++ {
			a11[i], a12[i], a21[i], a22[i] = a[i][:m], a[i][m:], a[i+m][:m], a[i+m][m:]
			b11[i], b12[i], b21[i], b22[i] = b[i][:m], b[i][m:], b[i+m][:m], b[i+m][m:]
		}

		// Realiza las operaciones del algoritmo de Strassen
		p1 := strassenNaiv(a11, sub(b12, b22))
		p2 := strassenNaiv(add(a11, a12), b22)
		p3 := strassenNaiv(add(a21, a22), b11)
		p4 := strassenNaiv(a22, sub(b21, b11))
		p5 := strassenNaiv(add(a11, a22), add(b11, b22))
		p6 := strassenNaiv(sub(a12, a22), add(b21, b22))
		p7 := strassenNaiv(sub(a11, a21), add(b11, b12))

		// Calcula los bloques de la matriz resultante
		c11 := add(sub(add(p5, p4), p2), p6)
		c12 := add(p1, p2)
		c21 := add(p3, p4)
		c22 := sub(sub(add(p5, p1), p3), p7)

		// Une los bloques para formar la matriz resultante
		for i := 0; i < m; i++ {
			copy(c[i][:m], c11[i])
			copy(c[i][m:], c12[i])
			copy(c[i+m][:m], c21[i])
			copy(c[i+m][m:], c22[i])
		}
	}

	return c
}

/*
Esta funcion es auxiliar del metodo strassenNaiv numero 9
*/
func add(a, b [][]int) [][]int {
	n := len(a)
	c := make([][]int, n)
	for i := range c {
		c[i] = make([]int, n)
		for j := 0; j < n; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

/*
Esta funcion es auxiliar del metodo strassenNaiv numero 9
*/
func sub(a, b [][]int) [][]int {
	n := len(a)
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
		for j := 0; j < n; j++ {
			c[i][j] = a[i][j] - b[i][j]
		}
	}
	return c
}

/*
Esta funcion es auxiliar del metodo strassenWinograd numero 10
*/
func addMatrix(a, b [][]int) [][]int {
	n := len(a)
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
		for j := 0; j < n; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

/*
Esta funcion es auxiliar del metodo strassenWinograd numero 10
*/
func subMatrix(a, b [][]int) [][]int {
	n := len(a)
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
		for j := 0; j < n; j++ {
			c[i][j] = a[i][j] - b[i][j]
		}
	}
	return c
}

/*
*
Esta funcion es auxiliar del metodo strassenWinograd numero 10
Función para dividir una matriz en cuatro submatrices más pequeñas
*
*/
func divideMatrix(m [][]int, half int) ([][]int, [][]int, [][]int, [][]int) {
	a := make([][]int, half)
	b := make([][]int, half)
	c := make([][]int, half)
	d := make([][]int, half)
	for i := 0; i < half; i++ {
		a[i] = make([]int, half)
		b[i] = make([]int, half)
		c[i] = make([]int, half)
		d[i] = make([]int, half)
		copy(a[i], m[i][:half])
		copy(b[i], m[i][half:])
		copy(c[i], m[i+half][:half])
		copy(d[i], m[i+half][half:])
	}
	return a, b, c, d
}

/*
*
Esta funcion pertenece al metodo strassenWinograd numero 10

*
*/
func strassenWinograd(a, b [][]int) [][]int {
	n := len(a)

	// Caso base
	if n == 1 {
		c := make([][]int, 1)
		c[0] = make([]int, 1)
		c[0][0] = a[0][0] * b[0][0]
		return c
	}

	// Dividir las matrices en submatrices más pequeñas
	half := n / 2
	a11, a12, a21, a22 := divideMatrix(a, half)
	b11, b12, b21, b22 := divideMatrix(b, half)

	// Calcular las submatrices recursivamente
	m1 := strassenWinograd(addMatrix(a11, a22), addMatrix(b11, b22))
	m2 := strassenWinograd(addMatrix(a21, a22), b11)
	m3 := strassenWinograd(a11, subMatrix(b12, b22))
	m4 := strassenWinograd(a22, subMatrix(b21, b11))
	m5 := strassenWinograd(addMatrix(a11, a12), b22)
	m6 := strassenWinograd(subMatrix(a21, a11), addMatrix(b11, b12))
	m7 := strassenWinograd(subMatrix(a12, a22), addMatrix(b21, b22))

	// Calcular las submatrices de la matriz resultado
	c11 := addMatrix(subMatrix(addMatrix(m1, m4), m5), m7)
	c12 := addMatrix(m3, m5)
	c21 := addMatrix(m2, m4)
	c22 := addMatrix(subMatrix(addMatrix(m1, m3), m2), m6)

	// Combinar las submatrices en la matriz resultado
	c := make([][]int, n)
	for i := 0; i < half; i++ {
		c[i] = make([]int, n)
		copy(c[i][:half], c11[i])
		copy(c[i][half:], c12[i])
	}
	for i := 0; i < half; i++ {
		c[i+half] = make([]int, n)
		copy(c[i+half][:half], c21[i])
		copy(c[i+half][half:], c22[i])
	}
	return c
}

/**

Este método aplica el metodo SequentialBlock numero 11
"bsize" es el tamaño de bloque utilizado para optimizar el rendimiento de la multiplicación de matrices.
**/

func SequentialBlock(A [][]int, B [][]int, size int, bsize int) [][]int {
	// Crear la matriz "C" de tamaño "size".
	C := make([][]int, size)
	for i := range C {
		// Inicializar cada fila de la matriz "C" con un slice de tamaño "size".
		C[i] = make([]int, size)
	}

	// Iterar sobre los bloques de tamaño "bsize" en las matrices "A", "B" y "C".
	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				// Iterar sobre cada entrada en los bloques actuales.
				for i := i1; i < i1+bsize && i < size; i++ {
					for j := j1; j < j1+bsize && j < size; j++ {
						for k := k1; k < k1+bsize && k < size; k++ {
							// Calcular la entrada correspondiente en la matriz resultante "C".
							C[i][j] += A[i][k] * B[k][j]
						}
					}
				}
			}
		}
	}

	// Devolver la matriz resultante "C".
	return C
}

/*
*

# Esta funcion aplica el metodo Parallel Block numero 12

*
*/
func matrixMultiplication(A, B, C [][]float64) [][]float64 {
	size := len(A)
	bsize := 32

	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				for i := i1; i < min(i1+bsize, size); i++ {
					for j := j1; j < min(j1+bsize, size); j++ {
						for k := k1; k < min(k1+bsize, size); k++ {
							A[k][i] += B[k][j] * C[j][i]
						}
					}
				}
			}
		}
	}

	return A
}

/*
*

# Esta funcion es auxiliar del metodo numero 11 Parallel Block

*
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
