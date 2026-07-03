package main

import (
	"fmt"
	"matrix/ex01/vector"
)

func main() {

	vectors := []vector.Vector[int]{
		{1, 2, 3},
		{2, 4, 7},
		{11, 2, 5},
	}

	coefs := []int{3, 2, 3}

	newVector := vector.LinearCombination(vectors, coefs)

	fmt.Printf("Combination: %v\n", newVector)
}
