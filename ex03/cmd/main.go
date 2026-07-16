package main

import (
	"fmt"
	"matrix/ex03/vector"
)

func main() {
	u := vector.Vector[float64]{0.0, 0.0}
	v := vector.Vector[float64]{1.0, 1.0}
	fmt.Printf("Dot 1: %f\n", u.Dot(v)) // expected: 0.0

	u = vector.Vector[float64]{1, 1}
	v = vector.Vector[float64]{1, 1}
	fmt.Printf("Dot 2: %f\n", u.Dot(v)) // expected: 2.0

	u = vector.Vector[float64]{-1, 6}
	v = vector.Vector[float64]{3, 2}
	fmt.Printf("Dot 3: %f\n", u.Dot(v)) // expected: 9.0

	u = vector.Vector[float64]{3, 2, 1}
	v = vector.Vector[float64]{5, -2, 0}
	fmt.Printf("Dot 4: %f\n", u.Dot(v)) // expected: 11.0

	u = vector.Vector[float64]{0, -2, 5, 1}
	v = vector.Vector[float64]{4, 1, -3, 2}
	fmt.Printf("Dot 5: %f\n", u.Dot(v)) // expected: -15.0

	u = vector.Vector[float64]{2.5, -1.5}
	v = vector.Vector[float64]{2.5, -1.5}
	fmt.Printf("Dot 6: %f\n", u.Dot(v)) // expected: 8.5
}
