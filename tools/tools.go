package tools

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"time"
)

const one_million = 1000 * 1000

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

func GetTmStmp() float64 { // Timestamp in milliseconds
	return float64(time.Now().UnixNano()) / one_million
}

func GetTmStmpUtc() float64 { // UTC timestamp in milliseconds
	return float64(time.Now().UTC().UnixNano()) / one_million
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
