package gcdtest

import (
	"math/rand"
	"testing"
	"time"

	"github.com/uscott/go-tools/tmath"
)

func TestGCD(t *testing.T) {
	a, b, c, d := 2, 3, 5, 7
	t.Logf("%d: %d\n", 1, tmath.GCD())
	t.Logf("%d: %d\n", 2, tmath.GCD(-a))
	t.Logf("%d: %d\n", 3, tmath.GCD(2, 3, 5))
	t.Logf("%d: %d\n", 4, tmath.GCD(a*b, b*c, b*d))
	x := []int{a, b, c}
	t.Logf("%d: %d\n", 5, tmath.GCD(x...))
	y := []int{a * b, b * c, b * 7}
	t.Logf("%d: %d\n", 6, tmath.GCD(y...))
	t.Logf("%d: %d\n", 7, tmath.GCD(a*b, b*c, a*b*c))
	t.Logf("%d: %d\n", 8, tmath.GCD(a*b, a*b*c, +a*a*c))
	t.Logf("%d: %d\n", 9, tmath.GCD(a*b, a*b*c, -a*a*c))
	rand.Seed(time.Now().UnixNano())
	p := []int{a, b, c}
	n := len(p)
	r := make([]int, 3)
	for i := range r {
		r[i] = 1
		for c := 0; c < 6; c++ {
			j := rand.Intn(n)
			q := p[j]
			r[i] *= q
		}
		t.Logf("r[%d] = %d\n", i, r[i])
	}
	gcd := tmath.GCD(r...)
	t.Logf("gcd = %d\n", gcd)
	for i, v := range r {
		t.Logf("r[%d]/gcd = %d\n", i, v/gcd)
	}
}
