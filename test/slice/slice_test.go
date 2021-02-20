package slicetest

import (
	"testing"

	"github.com/uscott/go-tools/slice"
)

type X struct {
	A int
	B string
}

func Test_Reverse(t *testing.T) {

	a := []int{1, 2, 3, 4}
	t.Logf("a = %+v\n", a)

	slice.Reverse(a)
	t.Logf("a = %+v\n", a)

	slice.Reverse(&a)
	t.Logf("a = %+v\n", a)

	b := [2]int{0, 1}
	t.Logf("b = %+v\n", b)

	slice.Reverse(&b)
	t.Logf("b = %+v\n", b)

	slice.RmByVal(&a, 1, 3)
	t.Logf("a = %+v\n", a)

	u := 5.2
	c := []float64{1.7, u}
	t.Logf("c = %+v\n", c)

	slice.RmByVal(&c, &u)
	t.Logf("c = %+v\n", c)

	x1, x2, x3 := X{1, "abc"}, X{20, "wat"}, X{-57, "hello"}
	vslc := []X{x1, x2, x3}
	t.Logf("vslc = %+v\n", vslc)

	slice.RmByVal(&vslc, &x1)
	t.Logf("vslc = %+v\n", vslc)
	slice.RmByVal(&vslc, x2)
	t.Logf("vslc = %+v\n", vslc)

	pslc := []*X{&x1, &x2, &x3}
	t.Log("pslc = ")
	for _, p := range pslc {
		t.Logf(" %+v ", *p)
	}

	slice.RmByVal(&pslc, &x2, &x3)
	t.Log("pslc = ")
	for _, p := range pslc {
		t.Logf(" %+v ", *p)
	}

	d := []string{"ab", "cd", "uvw", "xyz"}
	t.Logf("d = %+v\n", d)
	slice.RmByIndex(&d, 3, 0)
	t.Logf("d = %+v\n", d)
}

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
