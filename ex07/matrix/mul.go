package matrix

import "matrix/ex07/vector"

func (m Matrix[K]) MulMatrix(u Matrix[K]) Matrix[K] {

	if m.IsZero() || u.IsZero() || m.Size()[1] != u.Size()[0] {
		return nil
	}

	result := make(Matrix[K], m.Size()[0])

	for i := range m.Size()[0] {

		row := make([]K, u.Size()[1])

		for j := range u.Size()[1] {
			for k := range u.Size()[0] {
				row[j] = fma(m[i][k], u[k][j], row[j])
			}
		}

		result[i] = row
	}

	return result
}

func (m Matrix[K]) MulVector(v vector.Vector[K]) vector.Vector[K] {

	if m.IsZero() || v.IsZero() || m.Size()[1] != len(v) {
		return nil
	}

	result := make(vector.Vector[K], m.Size()[0])

	for i := range m.Size()[0] {

		for j := range m.Size()[1] {
			result[i] = fma(m[i][j], v[j], result[i])
		}
	}

	return result
}
