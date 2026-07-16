package matrix

func (m Matrix[K]) Inverse() (Matrix[K], error) {

	if m.IsZero() || !m.IsSquare() {
		return nil, nil
	}

	n := len(m)

	augmented := make(Matrix[K], n)
	for i := range n {
		row := make([]K, 2*n)
		copy(row, m[i])
		row[n+i] = 1
		augmented[i] = row
	}

	reduced := augmented.RowEchelon()

	inverse := make(Matrix[K], n)
	for i := range n {

		if reduced[i][i] != 1 {
			return nil, ErrMatrixSingular
		}

		for j := range n {
			if j != i && reduced[i][j] != 0 {
				return nil, ErrMatrixSingular
			}
		}

		row := make([]K, n)
		copy(row, reduced[i][n:])
		inverse[i] = row
	}

	return inverse, nil
}
