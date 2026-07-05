package main

import (
	"fmt"
	"matrix/ex04/vector"
)

func main() {
	u := vector.Vector[float64]{0, 0, 0}
	fmt.Printf("Norm 1: %f %f %f\n", u.Norm1(), u.Norm(), u.NormInf())

	u = vector.Vector[float64]{1, 2, 3}
	fmt.Printf("Norm 2: %f %f %f\n", u.Norm1(), u.Norm(), u.NormInf())

	u = vector.Vector[float64]{-1, -2}
	fmt.Printf("Norm 3: %f %f %f\n", u.Norm1(), u.Norm(), u.NormInf())
}