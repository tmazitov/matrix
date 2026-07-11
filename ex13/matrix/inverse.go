package matrix

func (m Matrix[K]) minor(x, y int) K {

	if len(m) == 1 {
		return m.Determinant()
	}

	matrix := make(Matrix[K], 0, len(m)-1)
	for k := range len(m) {
		if k == y {
			continue
		}
		row := make([]K, len(m))
		copy(row, m[k])
		row = append(row[:x], row[x+1:]...)
		matrix = append(matrix, row)
	}
	return matrix.Determinant()
}

func (m Matrix[K]) Inverse() Matrix[K] {

	if m.IsZero() || !m.IsSquare() {
		return nil
	}

	det := m.Determinant()
	if det == 0 {
		return nil
	}

	if len(m) == 1 {
		return Matrix[K]{{1 / m[0][0]}}
	}

	minorMatrix := make(Matrix[K], 0, len(m))
	for i := range len(m) {
		row := make([]K, len(m))
		for j := range len(m) {
			row[j] = m.minor(i, j)
			if (i+j)%2 == 1 && row[j] != 0 {
				row[j] *= -1
			}
		}
		minorMatrix = append(minorMatrix, row)
	}

	return minorMatrix.Scl(1 / det)
}
