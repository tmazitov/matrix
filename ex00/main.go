package main

import (
	"ex00/matrix"
	"ex00/vector"
	"fmt"
)

func main() {

	fmt.Println("Vector operations:")

	v := vector.Vector[int]{1, 2, 3}
	w := vector.Vector[int]{2, 3, 5}

	fmt.Printf("Add: %v\n", v.Add(w))
	fmt.Printf("Sub: %v\n", v.Sub(w))
	fmt.Printf("Scl: %v\n", v.Scl(5))

	a1 := matrix.Matrix[float64]{
		{1.1, 2.2, 3.5},
		{1.1, 2.2, 3.5},
		{1.1, 2.2, 3.5},
	}
	a2 := matrix.Matrix[float64]{
		{1.1, 2.2, 3.5},
		{1.1, 2.2, 3.5},
		{1.1, 2.2, 3.5},
	}
	fmt.Printf("Add: %v\n", a1.Add(a2))
	fmt.Printf("Sub: %v\n", a1.Sub(a2))
	fmt.Printf("Scl: %v\n", a1.Scl(10))
}
