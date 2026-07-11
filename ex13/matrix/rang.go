package matrix

func (m Matrix[K]) Rang() int {

	echelon := m.RowEchelon()

	rank := 0
	for _, row := range echelon {
		for _, value := range row {
			if value != 0 {
				rank++
				break
			}
		}
	}

	return rank
}