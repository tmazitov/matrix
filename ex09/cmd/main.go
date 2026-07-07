package main

import (
	"fmt"
	"matrix/ex09/matrix"
)

func main() {
	matrices := []matrix.Matrix[float64]{
		{
			{1, 2},
			{3, 4},
			{5, 6},
		},
		{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		},
		{
			{4, 2, 1},
			{-3, 6, 5},
			{8, -4, 0},
		},
	}

	for i, matrix := range matrices {
		fmt.Printf("Transpose %d\n", i)
		matrix.Transpose().PrettyPrint()
	}
}
