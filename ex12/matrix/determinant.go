package matrix

func (m Matrix[K]) submatrices() []Matrix[K] {
	result := make([]Matrix[K], len(m))

	for i := 0; i < len(m); i++ {
		matrix := make(Matrix[K], 0, len(m)-1)
		for k := 1; k < len(m); k++ {
			shadowedIndex := len(m) - i - 1
			row := make([]K, len(m))
			copy(row, m[k])
			row = append(row[:shadowedIndex], row[shadowedIndex+1:]...)
			matrix = append(matrix, row)
		}
		result[len(m)-i-1] = matrix
	}

	return result
}

func (m Matrix[K]) Determinant() K {
	size := m.Size()
	if m.IsZero() || !m.IsSquare() || size[0] > 4 {
		return 0
	}

	switch size[0] {
	case 1:
		return m[0][0]
	case 2:
		return m.determinant2()
	case 3:
		return m.determinant3()
	case 4:
		return m.determinant4()
	}

	return 1
}

func (m Matrix[K]) determinant4() K {

	subs := m.submatrices()

	var result K
	for i, matrix := range subs {
		scalar := K(1)
		if i%2 == 1 {
			scalar *= -1
		}
		result = fma(scalar*m[0][i], matrix.determinant3(), result)
	}
	return result
}

func (m Matrix[K]) determinant3() K {

	subs := m.submatrices()

	var result K
	for i, matrix := range subs {
		scalar := K(1)
		if i%2 == 1 {
			scalar *= -1
		}
		result = fma(scalar*m[0][i], matrix.determinant2(), result)
	}

	return result
}

func (m Matrix[K]) determinant2() K {
	return m[0][0]*m[1][1] - m[1][0]*m[0][1]
}
