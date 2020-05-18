package tools

import (
	"testing"
)

func TestPutGetMatrix(t *testing.T) {
	var u []float64
	u = make([]float64, 4)
	for i := 0; i < len(u); i++ {
		u[i] = float64(i)
	}
	var m [][]float64
	m = make([][]float64, 2)
	for i := 0; i < len(m); i++ {
		m[i] = make([]float64, 1)
	}
	PutMatrix(&u, &m, true)
	if m[0][0] != u[0] || m[0][1] != u[1] || m[1][0] != u[2] || m[0][1] != u[3] {
		t.Error("Something went wrong")
		return
	}
	var v []float64
	GetMatrix(&m, &v, true)
	if len(v) != len(u) || cap(v) != cap(u) {
		t.Error("Mismatched lengths")
		return
	}
	for i, _ := range u {
		if u[i] != v[i] {
			t.Error("Something went wrong")
			return
		}
	}
}
