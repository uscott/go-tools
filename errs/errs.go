package errs

import (
	"errors"
)

// Error values
var (
	ErrConversion = errors.New("Type Conversion Error")
	ErrKey        = errors.New("Map Key Error")
	ErrNilPtr     = errors.New("Nil Pointer")
)
