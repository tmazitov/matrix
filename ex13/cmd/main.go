package main

import (
	"fmt"

	"matrix/ex13/matrix"
)

func main() {

	matrices := []matrix.Matrix[float64]{
		// identity matrix -> full rank
		{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
		// linearly dependent rows -> rank 1
		{
			{1, 2},
			{2, 4},
		},
		// classic 42-subject example -> rank 3
		{
			{8, 5, -2},
			{4, 7, 20},
			{7, 6, 1},
		},
		// duplicate rows -> rank 1
		{
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3},
		},
		// all zeros -> rank 0
		{
			{0, 0, 0},
			{0, 0, 0},
		},
		// rank-deficient 4x4 (trivial zero row appended to a full-rank 3x3)
		{
			{2, 1, -1, 8},
			{-3, -1, 2, -11},
			{-2, 1, 2, -3},
			{0, 0, 0, 0},
		},
		// more columns than rows -> full row rank
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
		// 1x1
		{
			{5},
		},
	}

	for i, m := range matrices {
		fmt.Printf("== case %d ==\n", i+1)
		m.PrettyPrint()
		fmt.Println("rank:", m.Rang())
	}
}