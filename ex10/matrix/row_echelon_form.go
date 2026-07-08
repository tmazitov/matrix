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
			// fmt.Println("init pivot lvl", pivotLevel)
			nonZeroCells := []int{}

			for rowIndex, row := range rows {
				// fmt.Println(row)
				if len(row) > 0 && row[i] != 0 {
					nonZeroCells = append(nonZeroCells, rowIndex)
				}
			}
			// fmt.Printf("check column %d : %d\n", i, len(nonZeroCells))

			if len(nonZeroCells) > 1 && i > pivotLevel {
				pivot = rows[nonZeroCells[0]]
				pivotIndex = nonZeroCells[0]
				pivotLevel = i
				break
			}

			i++
		}

		// fmt.Printf("pivot : %v\n", pivot)

		if pivot == nil {
			break
		}

		result = append(result, pivot)
		for rowIndex, row := range rows {
			if len(row) == 0 || rowIndex == pivotIndex || row[pivotLevel] == 0 {
				continue
			}
			scale := row[pivotLevel] / pivot[pivotLevel]
			rows[rowIndex] = row.Add(pivot.Scl(-1 * scale))
			// fmt.Printf("row scale: %v %v -> %v\n", scale, row, rows[rowIndex])
		}

		// fmt.Printf("temp: %v\n", rows)

		if pivotIndex == 0 {
			rows = rows[1:]
		} else {
			rows = rows[pivotIndex-1 : pivotIndex]
		}
		pivot = nil
		pivotIndex = -1
	}

	for _, row := range rows {
		result = append(result, row)
	}

	return result
}
