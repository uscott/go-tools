package misc

import (
	"errors"
	"fmt"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
	"math"
	"time"
)

const one_million = 1000 * 1000

var ERR_CONVERSION = errors.New("CONVERSION ERROR")
var ERR_NIL_POINTER = errors.New("NIL POINTER")

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

func CopyMatrix(m *mat.Dense) *mat.Dense {
	if m == nil {
		return nil
	}
	nr, nc := m.Dims()
	return mat.NewDense(nr, nc, RawDataRef(m))
}

func CopyMap(m_in map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m_in {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}
	return cp
}

func GetRowVctr(src *[]float64, dst *[][]float64) error {
	switch {
	case src == nil:
		return fmt.Errorf("Nil source pointer")
	case dst == nil:
		return fmt.Errorf("Nil destination pointer")
	case len(*src) == 0:
		*dst = (*dst)[:0]
		return nil
	}
	if len(*dst) > 1 {
		*dst = (*dst)[:1]
	}
	return PutMatrix(src, dst, true)
}

func GetColVctr(src *[]float64, dst *[][]float64, byval bool) error {
	switch {
	case src == nil:
		return fmt.Errorf("Nil source pointer")
	case dst == nil:
		return fmt.Errorf("Nil destination pointer")
	case len(*src) == 0:
		*dst = (*dst)[:0]
		return nil
	}
	src_len := len(*src)
	if cap(*dst) < src_len {
		*dst = make([][]float64, src_len)
	}
	if len(*dst) != src_len {
		*dst = (*dst)[:src_len]
	}
	for i, x := range *src {
		p := &(*dst)[i]
		if byval {
			*p = make([]float64, 1)
			(*p)[0] = x
		} else {
			*p = (*src)[i : i+1]
		}
	}
	return nil
}

func GetTmStmp() float64 { // Timestamp in milliseconds
	return float64(time.Now().UnixNano()) / one_million
}

func GetTmStmpUtc() float64 { // UTC timestamp in milliseconds
	return float64(time.Now().UTC().UnixNano()) / one_million
}

func GetMatrix(src *[][]float64, dst *[]float64, row_major bool) error {
	switch {
	case src == nil:
		return fmt.Errorf("Nil source pointer")
	case dst == nil:
		return fmt.Errorf("Nil destination pointer")
	case len(*src) == 0:
		*dst = (*dst)[:0]
		return nil
	}
	var nrows, ncols int
	if row_major {
		nrows = len(*src)
		ncols = len((*src)[0])
	} else {
		ncols = len(*src)
		nrows = len((*src)[0])
	}
	if cap(*dst) < nrows*ncols {
		*dst = make([]float64, nrows*ncols)
	}
	if len(*dst) != nrows*ncols {
		*dst = (*dst)[:nrows*ncols]
	}
	switch row_major {
	case true:
		for i, r := range *src {
			if len(r) != ncols {
				return fmt.Errorf("Ragged matrix")
			}
			copy((*dst)[i*ncols:(i+1)*ncols], r)
		}
	case false:
		for j, c := range *src {
			if len(c) != nrows {
				return fmt.Errorf("Ragged matrix")
			}
			copy((*dst)[j*nrows:(j+1)*nrows], c)
		}
	}
	return nil
}

func PutMatrix(src *[]float64, dst *[][]float64, row_major bool) error {
	switch {
	case src == nil:
		return fmt.Errorf("Nil source pointer")
	case dst == nil:
		return fmt.Errorf("Nil destination pointer")
	case len(*src) == 0:
		*dst = (*dst)[:0]
		return nil
	}
	src_len := len(*src)
	if src_len%len(*dst) != 0 {
		return fmt.Errorf("Matrix dimension incompatible with slice source length")
	}
	var nrows, ncols int
	switch row_major {
	case true:
		nrows = len(*dst)
		ncols = src_len / nrows
		for i, r := range *dst {
			p := &(*dst)[i]
			if cap(r) < ncols {
				*p = make([]float64, ncols)
			}
			if len(r) != ncols {
				*p = (*p)[:ncols]
			}
			copy(*p, (*src)[i*ncols:(i+1)*ncols])
		}
	case false:
		ncols = len(*dst)
		nrows = src_len / ncols
		for j, c := range *dst {
			p := &(*dst)[j]
			if cap(c) < nrows {
				*p = make([]float64, nrows)
			}
			if len(c) != nrows {
				*p = (*p)[:nrows]
			}
			copy(*p, (*src)[j*nrows:(j+1)*nrows])
		}
	}
	return nil
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

func Icl(x float64) int {
	return int(math.Ceil(x))
}

func Ifl(x float64) int {
	return int(math.Floor(x))
}

func Ird(x float64) int {
	return int(math.Round(x))
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

func PrintMatrix(
	m *mat.Dense, transpose bool, scale float64, format string) {

	var ub1, ub2 int
	var x float64
	if m == nil {
		fmt.Println(m)
		return
	}
	if transpose {
		ub2, ub1 = m.Dims()
	} else {
		ub1, ub2 = m.Dims()
	}
	for i := 0; i < ub1; i++ {
		for j := 0; j < ub2; j++ {
			if transpose {
				x = m.At(j, i) * scale
			} else {
				x = m.At(i, j) * scale
			}
			fmt.Printf(format, x)
		}
		fmt.Println()
	}
}

func NowUtc() time.Time {
	return time.Now().UTC()
}

func Nrows(m *mat.Dense) int {
	if m == nil {
		return 0
	}
	nr, _ := m.Dims()
	return nr
}

func Ncols(m *mat.Dense) int {
	if m == nil {
		return 0
	}
	_, nc := m.Dims()
	return nc
}

func RawDataRef(m *mat.Dense) []float64 {
	if m == nil {
		return []float64{}
	}
	var g blas64.General
	g = m.RawMatrix()
	return g.Data
}

func SetCol(m *mat.Dense, col int, val float64) {
	for i := 0; i < Nrows(m); i++ {
		m.Set(i, col, val)
	}
}

func SetRow(m *mat.Dense, row int, val float64) {
	for j := 0; j < Ncols(m); j++ {
		m.Set(row, j, val)
	}
}
