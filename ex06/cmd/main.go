package main

import (
	"fmt"
	"matrix/ex06/vector"
)

func main() {
	u := vector.Vector[float64]{0, 0, 1}
	v := vector.Vector[float64]{1, 0, 0}
	fmt.Printf("Cross product 1: %v\n", vector.CrossProduct(u, v))

	u = vector.Vector[float64]{1, 2, 3}
	v = vector.Vector[float64]{4, 5, 6}
	fmt.Printf("Cross product 2: %v\n", vector.CrossProduct(u, v))

	u = vector.Vector[float64]{4, 2, -3}
	v = vector.Vector[float64]{-2, -5, 16}
	fmt.Printf("Cross product 3: %v\n", vector.CrossProduct(u, v))
}
