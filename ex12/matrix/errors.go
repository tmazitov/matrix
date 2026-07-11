package matrix

import "errors"

var (
	ErrMatrixNil           = errors.New("matrix error: initial values can't be nil")
	ErrMatrixInvalidColumn = errors.New("matrix error: initial values contains an invalid row")
	ErrMatrixSingular      = errors.New("matrix error: matrix is singular and has no inverse")
)
