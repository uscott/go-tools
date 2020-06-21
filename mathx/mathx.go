package mathx

import (
	"math"
	"time"
)

func ChunkCl(x, chk_sz float64) float64 {
	return chk_sz * math.Ceil(x/chk_sz)
}

func ChunkFl(x, chk_sz float64) float64 {
	return chk_sz * math.Floor(x/chk_sz)
}

func ChunkRd(x, chk_sz float64) float64 {
	return chk_sz * math.Round(x/chk_sz)
}

func Clamp(x, lb, ub float64) float64 {
	return math.Min(ub, math.Max(lb, x))
}

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

func Fdiv(numerator, denominator int) float64 {
	return float64(numerator) / float64(denominator)
}

func Fmax(args ...float64) (maxval float64) {
	if len(args) == 0 {
		maxval = math.Inf(-1)
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

func Fmin(args ...float64) (minval float64) {
	if len(args) == 0 {
		minval = math.Inf(1)
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

func Imax(args ...int) (maxval int) {
	if args == nil {
		maxval = math.MinInt64
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

func Imin(args ...int) (minval int) {
	if args == nil {
		minval = math.MaxInt64
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

// Tmax returns the maximum of its arguments
func Tmax(args ...time.Time) (maxval time.Time) {
	if args == nil {
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
	if args == nil {
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

func Integrate(f func(float64) float64, lb float64, ub float64, n uint) float64 {
	dx := (ub - lb) / float64(n)
	var x0, x1, val0, val1 float64
	x0, val0 = lb, f(lb)
	s := 0.0
	for i := 0; uint(i) < n; i++ {
		x1, val1 = x0+dx, f(x0+dx)
		s += 0.5 * (val0 + val1) * dx
		x0, val0 = x1, val1
	}
	return s
}

func Icl(x float64) int {
	return int(math.Ceil(x))
}

func Ifl(x float64) int {
	return int(math.Floor(x))
}

func Ird(x float64) int {
	return int(math.Round(x))
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
