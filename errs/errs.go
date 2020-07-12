package errs

import (
	"errors"
)

// Error values
var (
	ErrCapacity     = errors.New("insufficient capacity")
	ErrConversion   = errors.New("type conversion error")
	ErrDivByZero    = errors.New("divide by zero")
	ErrDimMismatch  = errors.New("dimension mismatch")
	ErrInf          = errors.New("+/-Infinity")
	ErrKey          = errors.New("map key error")
	ErrNaN          = errors.New("NaN")
	ErrNilPtr       = errors.New("nil pointer")
	ErrNotConnected = errors.New("not connected")
)
