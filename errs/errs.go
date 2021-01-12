package errs

import (
	"errors"
)

// Error values
var (
	Capacity     = errors.New("insufficient capacity")
	Conversion   = errors.New("type conversion error")
	DivByZero    = errors.New("divide by zero")
	DimMismatch  = errors.New("dimension mismatch")
	Inf          = errors.New("+/-Infinity")
	Key          = errors.New("map key error")
	NaN          = errors.New("NaN")
	NilPtr       = errors.New("nil pointer")
	NotConnected = errors.New("not connected")
)
