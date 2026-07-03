package number

type numberType interface {
	~int | ~float64
}

func Lerp[K numberType](v, u K, t float32) K {

	if t < 0 || t > 1 {
		return 0
	}

	return K(float32(u-v)*(t)) + v
}
