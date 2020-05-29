package errs

import (
	"errors"
)

// Error values
var (
	ErrCapacity    = errors.New("insufficient capacity")
	ErrConversion  = errors.New("type conversion error")
	ErrDivByZero   = errors.New("divide by zero")
	ErrDimMismatch = errors.New("dimension mismatch")
	ErrKey         = errors.New("map key error")
	ErrNilPtr      = errors.New("nil pointer")
)
