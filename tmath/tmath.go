package tmath

import (
	"fmt"
	"math"
	"sort"
	"time"
)

var (
	ceil  = math.Ceil
	floor = math.Floor
	inf   = math.Inf
	log10 = math.Log10
	min   = math.Min
	nan   = math.NaN
	round = math.Round
)

// Clamp is an alias for Fclamp
var (
	Clamp = Fclamp
)

// ChunkCl returns the least integer multiple of chunkSz
// greater than or equal to x
func ChunkCl(x, chunkSz float64) float64 {
	return chunkSz * ceil(x/chunkSz)
}

// ChunkFl returns the greatest integer multiple of chunkSz
// less than or equal to x
func ChunkFl(x, chunkSz float64) float64 {
	return chunkSz * floor(x/chunkSz)
}

// ChunkRd returns the integer multiple of chunkSz nearest x
func ChunkRd(x, chunkSz float64) float64 {
	return chunkSz * round(x/chunkSz)
}

// ChunkClToS returns the string representation of ChunkCl(x, chunkSz)
// with precision inferred from chunkSz
func ChunkClToS(x, chunkSz float64) string {
	if chunkSz <= 0 {
		return fmt.Sprintf("%v", nan())
	}
	prec := uint(ceil(-log10(min(1, chunkSz))))
	return FtoS(ChunkCl(x, chunkSz), prec)
}

// ChunkFlToS returns the string representation of ChunkFl(x, chunkSz)
// with precision inferred from chunkSz
func ChunkFlToS(x, chunkSz float64) string {
	if chunkSz <= 0 {
		return fmt.Sprintf("%v", nan())
	}
	prec := uint(ceil(-log10(min(1, chunkSz))))
	return FtoS(ChunkFl(x, chunkSz), prec)
}

// ChunkRdToS returns the string representation of ChunkRd(x, chunkSz)
// with precision inferred from chunkSz
func ChunkRdToS(x, chunkSz float64) string {
	if chunkSz <= 0 {
		return fmt.Sprintf("%v", nan())
	}
	prec := uint(ceil(-log10(min(1, chunkSz))))
	return FtoS(ChunkRd(x, chunkSz), prec)
}

// FtoS returns the string representation of x with the given precision
func FtoS(x float64, prec uint) string {
	return fmt.Sprintf(fmt.Sprintf("%%.%df", prec), x)
}

// Fclamp constrains x between lb and ub
func Fclamp(x, lb, ub float64) float64 {
	switch {
	case x < lb:
		return lb
	case ub < x:
		return ub
	default:
		return x
	}
}

// Iclamp constrains x between lb and ub
func Iclamp(x, lb, ub int) int {
	switch {
	case x < lb:
		return lb
	case ub < x:
		return ub
	default:
		return x
	}
}

// Fdiv returns the float64 division of integer arguments
func Fdiv(numerator, denominator int) float64 {
	return float64(numerator) / float64(denominator)
}

// Fmax returns the max of float64 arguments
func Fmax(args ...float64) (maxval float64) {
	if len(args) == 0 {
		maxval = inf(-1)
	} else {
		maxval = args[0]
		for _, x := range args[1:] {
			if x > maxval {
				maxval = x
			}
		}
	}
	return
}

// Fmin returns the min of float64 arguments
func Fmin(args ...float64) (minval float64) {
	if len(args) == 0 {
		minval = inf(1)
	} else {
		minval = args[0]
		for _, x := range args[1:] {
			if x < minval {
				minval = x
			}
		}
	}
	return
}

// Fsign returns +/- 1 or 0 according to the sign of the argument
func Fsign(x float64) float64 {
	if math.IsNaN(x) {
		return x
	} else if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

// GCD returns the greatest common divisor of its arguments
// Returns 0 if no arguments or at least one argument is zero
// Otherwise returns -1 if any are negative
func GCD(args ...int) (gcd int) {
	switch len(args) {
	case 0:
		gcd = 0
		return
	case 1:
		if args[0] == 0 {
			gcd = 0
		} else if args[0] < 0 {
			gcd = -1
		} else {
			gcd = args[0]
		}
		return
	default:
		for _, x := range args {
			if x == 0 {
				gcd = 0
				return
			} else if x < 0 {
				gcd = -1
			}
		}
		if gcd < 0 {
			return
		}
		ub := Imin(args...)
		for c := 1; c <= ub; c++ {
			ok := true
			for _, x := range args {
				if x%c != 0 {
					ok = false
					break
				}
			}
			if ok {
				gcd = c
			}
		}
		return
	}
}

// Imax returns the max of integer arguments
func Imax(args ...int) (maxval int) {
	if len(args) == 0 {
		maxval = math.MinInt64
		return
	}
	maxval = args[0]
	for _, x := range args[1:] {
		if x > maxval {
			maxval = x
		}
	}
	return
}

// Imin returns the minimum of integer arguments
func Imin(args ...int) (minval int) {
	if len(args) == 0 {
		minval = math.MaxInt64
		return
	}
	minval = args[0]
	for _, x := range args[1:] {
		if x < minval {
			minval = x
		}
	}
	return
}

// Isign is like Fsign but for Ints
func Isign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

// RoundToSigFigs rounds to the significant number of figures
// If sigfigs < 1 it assumes the minimal number of significant figures
// to the right of the decimal point
// If right == true that means sigfigs specifies the numbers of
// nonzero sigfigs to the right of the decimal point
func RoundToSigFigs(x float64, sigfigs int, right bool) float64 {

	if x == 0 {
		return 0
	}

	sign, x := Sign(x), math.Abs(x)

	var x1, x2 float64
	var n1, n2 int

	if sigfigs < 1 || right {
		x1, x2 = math.Modf(x)
	}
	if x1 > 0 {
		n1 = Ifl(log10(x1))
	}

	if sigfigs < 1 {
		if x2 > 0 {
			n2 = -Ifl(log10(x2))
		}
		sigfigs = n1 + n2 + 1
	} else if right {
		sigfigs += n1
	}

	n := Ifl(log10(x)) + 1 - sigfigs
	a := math.Pow10(n)

	return sign * a * round(x/a)
}

func Sign(x float64) float64 {
	return math.Copysign(1, x)
}

// Tmax returns the maximum of its arguments
func Tmax(args ...time.Time) (maxval time.Time) {
	if len(args) == 0 {
		maxval = time.Time{}
	} else {
		maxval = args[0]
		for _, t := range args[1:] {
			if t.After(maxval) {
				maxval = t
			}
		}
	}
	return
}

// Tmin returns the minimum of its arguments
func Tmin(args ...time.Time) (minval time.Time) {
	if len(args) == 0 {
		minval = time.Date(9999, 12, 31, 23, 59, 59, 0, time.Local)
	} else {
		minval = args[0]
		for _, t := range args[1:] {
			if t.Before(minval) {
				minval = t
			}
		}
	}
	return
}

// Trapezoidal implements the trapezoidal integration approximation without any panics
func Trapezoidal(x, f []float64) float64 {
	n := len(x)
	switch {
	case len(f) != n:
		return nan()
	case n < 2:
		return nan()
	case !sort.Float64sAreSorted(x):
		return nan()
	}
	integral := 0.0
	for i := 0; i < n-1; i++ {
		integral += 0.5 * (x[i+1] - x[i]) * (f[i+1] + f[i])
	}
	return integral
}

func Icl(x float64) int {
	return int(ceil(x))
}

func Ifl(x float64) int {
	return int(floor(x))
}

func Ird(x float64) int {
	return int(round(x))
}

func Ipow(x int, n uint) (y int) {
	if n == 0 {
		return 1
	}
	y = 1
	for n > 1 {
		if n%2 == 0 {
			x = x * x
			n = n / 2
		} else {
			y = x * y
			x = x * x
			n = (n - 1) / 2
		}
	}
	y *= x
	return
}
