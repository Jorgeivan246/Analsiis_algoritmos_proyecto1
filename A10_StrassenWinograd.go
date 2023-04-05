package main

type A10_StrassenWinograd struct {
}

func (s A10_StrassenWinograd) StrassenWinograd(a [][]int, b [][]int) [][]int {
	n := len(a)

	// Caso base
	if n == 1 {
		c := make([][]int, 1)
		c[0] = make([]int, 1)
		c[0][0] = a[0][0] * b[0][0]
		return c
	}

	aAux := make([][]int, n)
	bAux := make([][]int, n)

	for i := range aAux {
		aAux[i] = make([]int, n)
		copy(aAux[i], a[i])
		bAux[i] = make([]int, n)
		copy(bAux[i], b[i])
	}

	// Dividir las matrices en submatrices más pequeñas
	half := n / 2
	a11, a12, a21, a22 := s.divideMatrix(aAux, half)
	b11, b12, b21, b22 := s.divideMatrix(bAux, half)

	// Calcular las submatrices recursivamente
	p1 := s.StrassenWinograd(s.SumarMatrices(a11, a22), s.SumarMatrices(b11, b22))
	p2 := s.StrassenWinograd(s.SumarMatrices(a21, a22), b11)
	p3 := s.StrassenWinograd(a11, s.RestarMatrices(b12, b22))
	p4 := s.StrassenWinograd(a22, s.RestarMatrices(b21, b11))
	p5 := s.StrassenWinograd(s.SumarMatrices(a11, a12), b22)
	p6 := s.StrassenWinograd(s.RestarMatrices(a21, a11), s.SumarMatrices(b11, b12))
	p7 := s.StrassenWinograd(s.RestarMatrices(a12, a22), s.SumarMatrices(b21, b22))

	// Calcular las submatrices de la matriz resultado
	c11 := s.SumarMatrices(s.RestarMatrices(s.SumarMatrices(p1, p4), p5), p7)
	c12 := s.SumarMatrices(p3, p5)
	c21 := s.SumarMatrices(p2, p4)
	c22 := s.SumarMatrices(s.RestarMatrices(s.SumarMatrices(p1, p3), p2), p6)

	// Combinar las submatrices en la matriz resultado
	resultado := make([][]int, n)
	for i := 0; i < half; i++ {
		resultado[i] = make([]int, n)
		copy(resultado[i][:half], c11[i])
		copy(resultado[i][half:], c12[i])
	}
	for i := 0; i < half; i++ {
		resultado[i+half] = make([]int, n)
		copy(resultado[i+half][:half], c21[i])
		copy(resultado[i+half][half:], c22[i])
	}
	return resultado
}

/*
Esta funcion es auxiliar del metodo strassenWinograd numero 10
*/
func (s A10_StrassenWinograd) SumarMatrices(a [][]int, b [][]int) [][]int {
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
func (s A10_StrassenWinograd) RestarMatrices(a [][]int, b [][]int) [][]int {
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

func (s A10_StrassenWinograd) Run(a [][]int, b [][]int, matriz3 [][]int) [][]int {
	return s.StrassenWinograd(a, b)
}
