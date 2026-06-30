package combination

import (
	"math"
	"matrix/ex00/vector"
)

type numberType interface {
	~int | ~float64
}

func LinearCombination[K numberType](u []vector.Vector[K], coefs []K) vector.Vector[K] {

	if len(u) == 0 || len(u) != len(coefs) || !vector.IsEqualDimension(u...) {
		return nil
	}

	result := make(vector.Vector[K], len(u[0]))
	for i, vector := range u {
		for j, value := range vector {
			result[j] = K(math.FMA(float64(value), float64(coefs[i]), float64(result[j])))
		}
	}

	return result
}

func fma[K numberType](x, y, z K) K {
	return 0
}
