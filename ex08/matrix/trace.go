package matrix

func (m Matrix[K]) Trace() K {
	if m.IsZero() || !m.IsSquare() {
		return 0
	}

	result := K(0)

	for i := range m.Size()[0] {
		result += m[i][i]
	}

	return result
}
