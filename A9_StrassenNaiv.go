package main

type A9_StrassenNaiv struct {
}

/*
Metodo número 9
Esta funcion utiliza el metodo strassenNaiv para calcular la multiplicación
*/
func (s A9_StrassenNaiv) strassenNaiv(a [][] int, b [][]int) [][]int {
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
		p1 := s.strassenNaiv(a11, s.sub(b12, b22))
		p2 := s.strassenNaiv(s.add(a11, a12), b22)
		p3 := s.strassenNaiv(s.add(a21, a22), b11)
		p4 := s.strassenNaiv(a22, s.sub(b21, b11))
		p5 := s.strassenNaiv(s.add(a11, a22), s.add(b11, b22))
		p6 := s.strassenNaiv(s.sub(a12, a22), s.add(b21, b22))
		p7 := s.strassenNaiv(s.sub(a11, a21), s.add(b11, b12))

		// Calcula los bloques de la matriz resultante
		c11 := s.add(s.sub(s.add(p5, p4), p2), p6)
		c12 := s.add(p1, p2)
		c21 := s.add(p3, p4)
		c22 := s.sub(s.sub(s.add(p5, p1), p3), p7)

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
func (s A9_StrassenNaiv) add(a [][] int, b [][]int) [][]int {
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
func (s A9_StrassenNaiv) sub(a [][] int, b [][]int) [][]int {
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

func (s A9_StrassenNaiv) Run(a [][]int, b [][]int, matriz3 [][] int) [][]int {
	return s.strassenNaiv(a, b)
}
