package vector

import "math"

func (v Vector[K]) Norm1() float32 {

	var sum float32 = 0.0

	if len(v) == 0 {
		return 0
	}

	for _, value := range v {
		if value < 0 {
			value *= -1
		}
		sum += float32(value)
	}

	return sum
}

func (v Vector[K]) Norm() float32 {

	var sum float64 = 0.0

	if len(v) == 0 {
		return 0
	}

	for _, value := range v {
		sum += math.Pow(float64(value), 2)
	}

	return float32(math.Pow(sum, 0.5))
}

func (v Vector[K]) NormInf() float32 {

	if len(v) == 0 {
		return 0
	}

	max := float64(v[0])

	for _, value := range v {

		if value < 0 {
			value *= -1
		}

		max = math.Max(max, float64(value))
	}
	return float32(max)
}
