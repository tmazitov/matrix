package vector

type numberType interface {
	~int | ~int64 | ~float64
}

type Vector[K numberType] []K

func NewVector[K numberType](values []K) (Vector[K], error) {
	if len(values) == 0 {
		return nil, ErrVectorNil
	}
	return Vector[K](values), nil
}
