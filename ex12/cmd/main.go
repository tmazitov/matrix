package main

import (
	"fmt"
	"matrix/ex12/matrix"
)

func main() {

	matrices := []matrix.Matrix[float64]{
		{
			{-8},
		},
		{
			{2, 3},
			{1, -2},
		},
		{
			{2, 0, 0},
			{0, 2, 0},
			{0, 0, 2},
		},
		{
			{8., 5., -2.},
			{4., 7., 20.},
			{7., 6., 1.},
		},
	}

	for i, matrix := range matrices {
		fmt.Printf("%d | For matrix :\n", i)
		matrix.PrettyPrint()

		fmt.Println("Inverse version is:")
		v := matrix.Inverse()
		v.PrettyPrint()

		fmt.Println("Check (A * A-1):")
		matrix.MulMatrix(v).PrettyPrint()
		fmt.Println("Check (A-1 * A):")
		v.MulMatrix(matrix).PrettyPrint()
	}
}
