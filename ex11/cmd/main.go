package main

import (
	"fmt"

	"matrix/ex11/matrix"
)

func main() {

	m2 := matrix.Matrix[float64]{
		{1, -1},
		{2, 3},
	}

	m3 := matrix.Matrix[float64]{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	m4 := matrix.Matrix[float64]{
		{8, 5, -2, 4},
		{4, 2.5, 20, 4},
		{8, 5, 1, 4},
		{28, -4, 17, 1},
	}

	fmt.Println("== 2x2 ==")
	m2.PrettyPrint()
	fmt.Println("determinant:", m2.Determinant())

	fmt.Println("== 3x3 ==")
	m3.PrettyPrint()
	fmt.Println("determinant:", m3.Determinant())

	fmt.Println("== 4x4 ==")
	m4.PrettyPrint()
	fmt.Println("determinant:", m4.Determinant())
}
