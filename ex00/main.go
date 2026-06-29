package main

import (
	"ex00/vector"
	"fmt"
)

func main() {

	v := vector.Vector[int]{1, 2, 3}
	w := vector.Vector[int]{2, 3, 5}

	fmt.Printf("Sum: %v\n", v.Sum(w))
	fmt.Printf("Sub: %v\n", v.Sub(w))
	fmt.Printf("Scl: %v\n", v.Scl(5))
}
