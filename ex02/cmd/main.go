package main

import (
	"fmt"
	"matrix/ex02/matrix"
	"matrix/ex02/number"
	"matrix/ex02/vector"
)

func main() {

	fmt.Printf("Numbers' lerp: %v\n", number.Lerp(0.0, 1.0, 0.5))

	v := vector.Vector[float64]{2, 1}
	u := vector.Vector[float64]{4, 2}
	fmt.Printf("Vectors' lerp : %f\n", vector.Lerp(v, u, 0.3))

	a := matrix.Matrix[float64]{
		{2, 1},
		{3, 4},
	}
	b := matrix.Matrix[float64]{
		{20, 10},
		{30, 40},
	}
	fmt.Printf("Matrices' lerp : %f\n", matrix.Lerp(a, b, 0.5))
}
