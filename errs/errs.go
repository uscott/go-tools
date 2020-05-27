package errs

import (
	"errors"
)

// Error values
var (
	ErrCapacity    = errors.New("Insufficient Capacity")
	ErrConversion  = errors.New("Type Conversion Error")
	ErrDimMismatch = errors.New("Dimension Mismatch")
	ErrKey         = errors.New("Map Key Error")
	ErrNilPtr      = errors.New("Nil Pointer")
)
