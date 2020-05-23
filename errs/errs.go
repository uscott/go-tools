package errs

import (
	"errors"
)

var Conversion = errors.New("Type Conversion Error")
var NilPtr = errors.New("Nil Pointer")
