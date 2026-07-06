package main

import (
	"fmt"
	"matrix/ex07/matrix"
	"matrix/ex07/vector"
)

func main() {
	v := vector.Vector[float64]{4, 2, 1}

	matrices := []matrix.Matrix[float64]{
		{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
		{
			{2, 0, 0},
			{0, 2, 0},
			{0, 0, 2},
		},
		{
			{2, -2, 2},
			{-2, 2, -2},
			{2, -2, 2},
		},
		{
			{2, -1, 3},
			{0, 4, -2},
			{1, 5, -3},
		},
	}

	for i, matrix := range matrices {
		fmt.Printf("Mul Matrix x Vector %d: %v\n", i+1, matrix.MulVector(v))
	}

	fmt.Printf("\n-------------------\n\n")

	u := matrix.Matrix[float64]{
		{3, -5, 2},
		{6, 8, -4},
		{7, 1, -3},
	}

	for i, matrix := range matrices {
		fmt.Printf("Mul Matrix x Matrix %d: %v\n", i+1, matrix.MulMatrix(u))
	}

	a := matrix.Matrix[float64]{
		{1, 2, 3},
		{4, 5, 6},
	}
	b := matrix.Matrix[float64]{
		{7, 8, 9, 10},
		{11, 12, 13, 14},
		{15, 16, 17, 18},
	}

	fmt.Printf("Mul Matrix x Matrix (diff size): %v\n", a.MulMatrix(b))
}
