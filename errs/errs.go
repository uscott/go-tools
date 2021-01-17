package errs

import (
	"errors"
)

// Error values
var (
	Capacity         = errors.New("Insufficient capacity")
	Conversion       = errors.New("Type conversion error")
	DivByZero        = errors.New("Divide by zero")
	DimMismatch      = errors.New("Dimension mismatch")
	Inf              = errors.New("+/-Infinity")
	Key              = errors.New("Map key error")
	NaN              = errors.New("NaN")
	NilPtr           = errors.New("Nil pointer")
	NilPtrArg        = errors.New("Nil input argument pointer")
	NilPtrReturn     = errors.New("Nil pointer return value")
	NilPtrUnexpected = errors.New("Unexpected nil pointer")
	NotConnected     = errors.New("Not connected")
)
