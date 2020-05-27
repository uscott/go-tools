package errs

import (
	"errors"
)

// Error values
var (
	ErrConversion = errors.New("Type Conversion Error")
	ErrNilPtr     = errors.New("Nil Pointer")
)
