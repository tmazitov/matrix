package matrix

type numberType interface {
	~int | ~float64
}

type Matrix[K numberType] [][]K

func (m Matrix[K]) Size() [2]int {

	M := len(m)
	if M == 0 {
		return [2]int{0, 0}
	}

	N := len(m[0])
	return [2]int{M, N}
}

func (m Matrix[K]) IsZero() bool {
	size := m.Size()
	return size[0] == 0 || size[1] == 0
}

func (m Matrix[K]) IsSquare() bool {
	size := m.Size()
	return size[0] == size[1]
}

func (m Matrix[K]) GetValue(i, j int) (K, bool) {

	size := m.Size()

	if i < 0 || i >= size[0] ||
		j < 0 || j >= size[1] {
		return 0, false
	}

	return m[i][j], true
}

func (m Matrix[K]) IsEqualSize(size [2]int) bool {

	s := m.Size()

	return s[0] == size[0] &&
		s[1] == size[1]
}
