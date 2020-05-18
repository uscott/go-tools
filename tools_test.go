package tools

import (
	"fmt"
	"testing"
)

func Test_PutGetMatrix(t *testing.T) {
	var u, v []float64
	u = make([]float64, 4)
	v = make([]float64, 4)
	for i := 0; i < len(u); i++ {
		u[i] = float64(i)
		v[i] = u[i]
	}
	fmt.Println(u)
	var m [][]float64
	m = make([][]float64, 2)
	for i := 0; i < len(m); i++ {
		m[i] = make([]float64, 1)
	}
	PutMatrix(&u, &m, true)
	fmt.Println(u)
	for _, r := range m {
		for _, x := range r {
			fmt.Printf("%v ", x)
		}
		fmt.Println()
	}
	if m[0][0] != u[0] || m[0][1] != u[1] || m[1][0] != u[2] || m[1][1] != u[3] {
		t.Error("Something went wrong")
		return
	}
	for i, x := range u {
		if x != v[i] {
			t.Error("Wrong")
		}
	}
	GetMatrix(&m, &v, true)
	fmt.Println(v)
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
	m[0][0] = 20
	if u[0] == 20 {
		t.Error()
	}
	for _, r := range m {
		for _, x := range r {
			fmt.Printf("%v ", x)
		}
		fmt.Println()
	}
	fmt.Println(u)
	fmt.Println(v)
}

func Test_GetRowVctr(t *testing.T) {
	var v = make([]float64, 4)
	v[0] = 1.0
	v[1] = 3.0
	fmt.Println(v)
	var rv = make([][]float64, 5)
	var err error
	err = GetRowVctr(&v, &rv)
	if err != nil {
		t.Error(err)
	}
	for _, r := range rv {
		for _, x := range r {
			fmt.Printf("%v ", x)
		}
		fmt.Println()
	}
	if len(rv) != 1 {
		t.Error("Wrong no. of rows")
	}
	for _, r := range rv {
		if len(r) != len(v) {
			t.Error("Wrong length")
		}
	}
	if v[0] != rv[0][0] || v[1] != rv[0][1] || v[2] != rv[0][2] || v[3] != rv[0][3] {
		t.Error("Not equal")
	}
}

func Test_GetColVctr(t *testing.T) {
	vlen := 1000
	var v = make([]float64, vlen)
	v[0] = 1
	v[1] = 2
	v[2] = 4
	var cv_val, cv_ref = make([][]float64, 2), make([][]float64, 2)
	var err error
	if err = GetColVctr(&v, &cv_val, true); err != nil {
		t.Error(err)
	}
	if err = GetColVctr(&v, &cv_ref, false); err != nil {
		t.Error(err)
	}
	if len(cv_val) != len(v) || len(cv_ref) != len(v) {
		t.Error()
	}
	for i, x := range v {
		if x != cv_val[i][0] || x != cv_ref[i][0] {
			t.Error()
		}
	}
	cv_ref[0][0] = 10
	if v[0] != 10 {
		t.Error()
	}
	if cv_val[0][0] == 10 {
		t.Error()
	}
}
