package matrix

func NewMatrix[K numberType](rows [][]K) (Matrix[K], error) {

	if rows == nil {
		return nil, ErrMatrixNil
	}

	if len(rows) > 0 {
		width := len(rows[0])
		for _, row := range rows {
			if len(row) != width {
				return nil, ErrMatrixInvalidColumn
			}
		}
	}

	return Matrix[K](rows), nil
}
