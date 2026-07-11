package matrix

func (m Matrix[K]) Transpose() Matrix[K] {
	if m.IsZero() {
		return Matrix[K]{}
	}

	size := m.Size()
	result := make(Matrix[K], size[1])

	for i := range size[0] {

		for j := range size[1] {
			if len(result[j]) == 0 {
				result[j] = make([]K, size[0])
			}
			result[j][i] = m[i][j]
		}
	}

	return result
}
