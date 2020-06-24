package slice

import (
	"testing"
)

func TestRev(t *testing.T) {
	n := 3
	x := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = float64(i)
	}
	y := make([]float64, n)
	copy(y, x)
	FltRev(y)
	for i := 0; i < n; i++ {
		if x[i] != y[n-i-1] {
			t.FailNow()
		}
	}
}
