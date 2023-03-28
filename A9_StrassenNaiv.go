package main

type A9_StrassenNaiv struct {
}

func (s A9_StrassenNaiv) StrassenNaiv(a [][]int, b [][]int) [][]int {
	size := len(a)
	resultado := make([][]int, size)

	for i := range resultado {
		resultado[i] = make([]int, size)
	}

	if size == 1 {
		return [][]int{{a[0][0] * b[0][0]}}
	} else {
		// Divide las matrices en bloques más pequeños
		m := size / 2
		a11 := make([][]int, m)
		a12 := make([][]int, m)
		a21 := make([][]int, m)
		a22 := make([][]int, m)

		b11 := make([][]int, m)
		b12 := make([][]int, m)
		b21 := make([][]int, m)
		b22 := make([][]int, m)

		for i := 0; i < m; i++ {
			a11[i] = a[i][:m]
			a12[i] = a[i][m:]
			a21[i] = a[m+i][:m]
			a22[i] = a[m+i][m:]

			b11[i] = b[i][:m]
			b12[i] = b[i][m:]
			b21[i] = b[m+i][:m]
			b22[i] = b[m+i][m:]
		}

		// Realiza las operaciones del algoritmo de Strassen
		p1 := s.StrassenNaiv(s.SumarMatrices(a11, a22), s.SumarMatrices(b11, b22))
		p2 := s.StrassenNaiv(s.SumarMatrices(a21, a22), b11)
		p3 := s.StrassenNaiv(a11, s.RestarMatrices(b12, b22))
		p4 := s.StrassenNaiv(a22, s.RestarMatrices(b21, b11))
		p5 := s.StrassenNaiv(s.SumarMatrices(a11, a12), b22)
		p6 := s.StrassenNaiv(s.RestarMatrices(a21, a11), s.SumarMatrices(b11, b12))
		p7 := s.StrassenNaiv(s.RestarMatrices(a12, a22), s.SumarMatrices(b21, b22))

		// Calcula los bloques de la matriz resultante
		c11 := s.SumarMatrices(s.RestarMatrices(s.SumarMatrices(p1, p4), p5), p7)
		c12 := s.SumarMatrices(p3, p5)
		c21 := s.SumarMatrices(p2, p4)
		c22 := s.SumarMatrices(s.RestarMatrices(s.SumarMatrices(p1, p3), p2), p6)

		// Une los bloques en una sola matriz
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				resultado[i][j] = c11[i][j]
				resultado[i][j+m] = c12[i][j]
				resultado[i+m][j] = c21[i][j]
				resultado[i+m][j+m] = c22[i][j]
			}
		}
	}

	return resultado
}

func (s A9_StrassenNaiv) SumarMatrices(a [][]int, b [][]int) [][]int {
	size := len(a)
	resultado := make([][]int, size)
	for i := range resultado {
		resultado[i] = make([]int, size)
		for j := 0; j < size; j++ {
			resultado[i][j] = a[i][j] + b[i][j]
		}
	}
	return resultado
}

func (s A9_StrassenNaiv) RestarMatrices(a [][]int, b [][]int) [][]int {
	size := len(a)
	resultado := make([][]int, size)
	for i := 0; i < size; i++ {
		resultado[i] = make([]int, size)
		for j := 0; j < size; j++ {
			resultado[i][j] = a[i][j] - b[i][j]
		}
	}
	return resultado
}

func (s A9_StrassenNaiv) Run(a [][]int, b [][]int, matriz3 [][]int) [][]int {
	return s.StrassenNaiv(a, b)
}
