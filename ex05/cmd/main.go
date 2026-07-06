package main

import (
	"fmt"
	"matrix/ex05/vector"
)

func main() {
	u := vector.Vector[float64]{1., 0.}
	v := vector.Vector[float64]{1., 0.}
	fmt.Println(u.AngleCos(v))
	// 1.0

	u = vector.Vector[float64]{1., 0.}
	v = vector.Vector[float64]{0., 1.}
	fmt.Println(u.AngleCos(v))
	// 0.0

	u = vector.Vector[float64]{-1., 1.}
	v = vector.Vector[float64]{1., -1.}
	fmt.Println(u.AngleCos(v))
	// -1.0

	u = vector.Vector[float64]{2., 1.}
	v = vector.Vector[float64]{4., 2.}
	fmt.Println(u.AngleCos(v))
	// 1.0

	u = vector.Vector[float64]{1., 2., 3.}
	v = vector.Vector[float64]{4., 5., 6.}
	fmt.Println(u.AngleCos(v))
	// 0.974631846
}
