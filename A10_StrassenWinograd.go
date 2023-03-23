package main

type A10_StrassenWinograd struct {
}

func (s A10_StrassenWinograd) strassenWinograd(a [][] int, b [][]int) [][]int {
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
	a11, a12, a21, a22 := s.divideMatrix(a, half)
	b11, b12, b21, b22 := s.divideMatrix(b, half)

	// Calcular las submatrices recursivamente
	m1 := s.strassenWinograd(s.addMatrix(a11, a22), s.addMatrix(b11, b22))
	m2 := s.strassenWinograd(s.addMatrix(a21, a22), b11)
	m3 := s.strassenWinograd(a11, s.subMatrix(b12, b22))
	m4 := s.strassenWinograd(a22, s.subMatrix(b21, b11))
	m5 := s.strassenWinograd(s.addMatrix(a11, a12), b22)
	m6 := s.strassenWinograd(s.subMatrix(a21, a11), s.addMatrix(b11, b12))
	m7 := s.strassenWinograd(s.subMatrix(a12, a22), s.addMatrix(b21, b22))

	// Calcular las submatrices de la matriz resultado
	c11 := s.addMatrix(s.subMatrix(s.addMatrix(m1, m4), m5), m7)
	c12 := s.addMatrix(m3, m5)
	c21 := s.addMatrix(m2, m4)
	c22 := s.addMatrix(s.subMatrix(s.addMatrix(m1, m3), m2), m6)

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

/*
Esta funcion es auxiliar del metodo strassenWinograd numero 10
*/
func (s A10_StrassenWinograd) addMatrix(a [][] int, b [][]int) [][]int {
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
func (s A10_StrassenWinograd) subMatrix(a [][] int, b [][]int) [][]int {
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
func (s A10_StrassenWinograd) divideMatrix(m [][]int, half int) ([][]int, [][]int, [][]int, [][]int) {
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

func (s A10_StrassenWinograd) Run(a [][]int, b [][]int, matriz3 [][] int) [][]int {
	return s.strassenWinograd(a, b)
}
