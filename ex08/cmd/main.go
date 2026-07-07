package main

import (
	"fmt"
	"matrix/ex08/matrix"
)

func main() {
	matrices := []matrix.Matrix[float64]{
		{
			{1, 0},
			{0, 1},
		},
		{
			{2., -5., 0.},
			{4., 3., 7.},
			{-2., 3., 4.},
		},
		{
			{-2., -8., 4.},
			{1., -23., 4.},
			{0., 6., 4.},
		},
	}
	
	for i, matrix := range matrices {
		fmt.Printf("Trace %d: %f\n", i+1, matrix.Trace())
	}
}

