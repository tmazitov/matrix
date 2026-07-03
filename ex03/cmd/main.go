package main

import (
	"fmt"
	"matrix/ex03/vector"
)

func main() {
	u := vector.Vector[float64]{0.0, 0.0}
	v := vector.Vector[float64]{1.0, 1.0}
	fmt.Printf("Dot 1: %f\n", u.Dot(v))

	u = vector.Vector[float64]{1, 1}
	v = vector.Vector[float64]{1, 1}
	fmt.Printf("Dot 2: %f\n", u.Dot(v))

	u = vector.Vector[float64]{-1, 6}
	v = vector.Vector[float64]{3, 2}
	fmt.Printf("Dot 3: %f\n", u.Dot(v))
}
