package vector

type numberType interface {
	~int | ~int64 | ~float64
}

type Vector[K numberType] []K
