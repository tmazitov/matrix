package matrix

import "math"

func fma[K numberType](x, y, z K) K {
	return K(math.FMA(float64(x), float64(y), float64(z)))
}

func Lerp[K numberType](v, u Matrix[K], t float32) Matrix[K] {

	if !u.IsEqualSize(v.Size()) || t < 0 || t > 1 {
		return nil
	}

	result := make(Matrix[K], len(v))
	for i := range v.Size()[0] {

		temp := make([]K, v.Size()[1])

		for j := range v.Size()[1] {
			temp[j] = fma(u[i][j]-v[i][j], K(t), v[i][j])
		}

		result[i] = temp
	}
	return result

}
