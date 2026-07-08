package main

import (
	"fmt"

	"matrix/ex10/matrix"
)

func main() {

	matrices := []matrix.Matrix[float64]{
		// identity matrix
		{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
		// simple 2x2
		{
			{1, 2},
			{3, 4},
		},
		// linearly dependent rows -> zero row expected
		{
			{1, 2},
			{2, 4},
		},
		// classic 42-subject example
		// expected: [1, 0.625, 0, 0, -12.1666667]
		//           [0, 0, 1, 0, -3.6666667]
		//           [0, 0, 0, 1, 29.5]
		{
			{8, 5, -2, 4, 28},
			{4, 2.5, 20, 4, -4},
			{8, 5, 1, 4, 17},
		},
		// original 3x3
		{
			{1, 1, 2},
			{1, 2, 3},
			{3, 4, 5},
		},
		// zero pivot requiring a row swap
		{
			{0, 1, 1},
			{2, 4, 6},
			{1, 3, 5},
		},
		// zero row in the middle
		{
			{1, 2, 3},
			{0, 0, 0},
			{4, 5, 6},
		},
		// more rows than columns
		{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
		},
		// more columns than rows
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
		// single row
		{
			{2, 4, 6},
		},
		// single column
		{
			{2},
			{4},
			{6},
		},
		// 1x1
		{
			{5},
		},
		// all zeros
		{
			{0, 0, 0},
			{0, 0, 0},
		},
		// negative and fractional values
		{
			{-1, 2.5, -3},
			{4, -5, 6.5},
			{-7.5, 8, 9},
		},
		// zero pivot column followed by usable pivot further right
		// (whole first column is zero -> should skip to the second column)
		{
			{0, 1, 2},
			{0, 3, 4},
			{0, 5, 7},
		},
		// duplicate rows -> rank 1
		{
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3},
		},
		// already in reduced row echelon form -> should be unchanged
		{
			{1, 0, 0, 5},
			{0, 1, 0, -3},
			{0, 0, 1, 2},
		},
		// rank-deficient 4x4 (last row is a trivial all-zero row appended
		// to an otherwise full-rank 3x3 system)
		{
			{2, 1, -1, 8},
			{-3, -1, 2, -11},
			{-2, 1, 2, -3},
			{0, 0, 0, 0},
		},
		// reverse-identity permutation, needs multiple row swaps
		// (row0<->row3 to fix column 0, then row1<->row2 to fix column 1)
		{
			{0, 0, 0, 1},
			{0, 0, 1, 0},
			{0, 1, 0, 0},
			{1, 0, 0, 0},
		},
		// very small / very large magnitude values
		{
			{1e-8, 2, 3},
			{4, 5e8, 6},
			{7, 8, 9e-3},
		},
	}

	for i, m := range matrices {
		fmt.Printf("== case %d ==\n", i+1)
		m.RowEchelon().PrettyPrint()
	}
}
