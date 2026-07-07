package matrix

func (m Matrix[K]) Add(v Matrix[K]) Matrix[K] {

	if !m.IsEqualSize(v.Size()) {
		return nil
	}

	sumOperation := func(i, j int) K {
		valueM := m[i][j]
		valueV := v[i][j]

		return valueM + valueV
	}

	return m.aggregate(sumOperation)
}

func (m Matrix[K]) Sub(v Matrix[K]) Matrix[K] {

	if !m.IsEqualSize(v.Size()) {
		return nil
	}

	subOperation := func(i, j int) K {

		valueM := m[i][j]
		valueV := v[i][j]

		return valueM - valueV
	}

	return m.aggregate(subOperation)
}

func (m Matrix[K]) Scl(a K) Matrix[K] {

	sclOperation := func(i, j int) K {

		valueM := m[i][j]

		return valueM * a
	}

	return m.aggregate(sclOperation)
}

type aggregateOperation[K numberType] func(i, j int) K

func (m Matrix[K]) aggregate(operation aggregateOperation[K]) Matrix[K] {

	size := m.Size()

	values := make([][]K, 0, size[0])

	for i, column := range m {

		newColumn := make([]K, 0, size[1])

		for j := range column {
			newColumn = append(newColumn, operation(i, j))
		}

		values = append(values, newColumn)
	}

	return (Matrix[K])(values)
}
