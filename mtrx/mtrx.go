package mtrx

import (
	"fmt"
	"github.com/uscott/gotools/errs"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
)

func CopyMatrix(m *mat.Dense) *mat.Dense {
	if m == nil {
		return nil
	}
	nr, nc := m.Dims()
	return mat.NewDense(nr, nc, RawDataRef(m))
}

func GetRowVctr(src *[]float64, dst *[][]float64) error {
	switch {
	case src == nil || dst == nil:
		return errs.NilPtr
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
	case src == nil || dst == nil:
		return errs.NilPtr
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

func GetMatrix(src *[][]float64, dst *[]float64, row_major bool) error {
	switch {
	case src == nil:
		return errs.NilPtr
	case dst == nil:
		return errs.NilPtr
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

func PutMatrix(src *[]float64, dst *[][]float64, row_major bool) error {
	switch {
	case src == nil:
		return errs.NilPtr
	case dst == nil:
		return errs.NilPtr
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
