package matrix

import (
	"matrix/ex10/vector"
)

func (m Matrix[K]) RowEchelon() Matrix[K] {

	result := make(Matrix[K], 0, len(m))

	rows := make([]vector.Vector[K], 0, len(m))
	for _, row := range m {
		rows = append(rows, vector.Vector[K](row))
	}

	var pivot vector.Vector[K]
	var pivotLevel int = -1
	var pivotIndex int = -1
	for {

		i := pivotLevel + 1

		for i < m.Size()[1] {
			nonZeroCells := []int{}

			for rowIndex, row := range rows {
				if len(row) > 0 && row[i] != 0 {
					nonZeroCells = append(nonZeroCells, rowIndex)
				}
			}

			if len(nonZeroCells) > 0 && i > pivotLevel {
				pivot = rows[nonZeroCells[0]]
				pivotIndex = nonZeroCells[0]
				pivotLevel = i
				break
			}

			i++
		}

		if pivot == nil {
			break
		}

		pivot = pivot.Scl(1 / pivot[pivotLevel])
		rows[pivotIndex] = pivot

		result = append(result, pivot)
		for rowIndex, r := range result[:len(result)-1] {
			row := vector.Vector[K](r)
			if row[pivotLevel] == 0 {
				continue
			}
			scale := row[pivotLevel] / pivot[pivotLevel]
			result[rowIndex] = row.Add(pivot.Scl(-1 * scale))
		}
		for rowIndex, row := range rows {
			if len(row) == 0 || rowIndex == pivotIndex || row[pivotLevel] == 0 {
				continue
			}
			scale := row[pivotLevel] / pivot[pivotLevel]
			rows[rowIndex] = row.Add(pivot.Scl(-1 * scale))
		}

		rows = append(rows[:pivotIndex], rows[pivotIndex+1:]...)
		pivot = nil
		pivotIndex = -1
	}

	for _, row := range rows {
		result = append(result, row)
	}

	return result
}
