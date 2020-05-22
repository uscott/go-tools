package rndm

import (
	"github.com/uscott/gotools/misc"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/mat"
	// "math/rand"
	// "time"
)

type RNG struct {
	*rand.Rand
	Source 	*rand.LockedSource
}

func (rng RNG) Seed(seed uint64) {
	rng.Source.Seed(seed)
}

func (rng RNG) Uint64() uint64 {
	return rng.Rand.Uint64()
}

// func NewRngSeeded(s uint64) RNG {
// 	p.Seed(s + uint64(time.Now().UnixNano()))
// }

func NewRng() RNG {
	var src = rand.LockedSource{}
	var p = &src
	return RNG{rand.New(p), p}
}

func NormRand(r RNG, chol *mat.TriDense, eps []float64) error {
	if chol == nil || r.Rand == nil || eps == nil {
		return misc.ERR_NIL_POINTER
	}
	n, _ := chol.Dims()
	u := mat.NewDense(n, 1, eps)
	slc := u.RawMatrix().Data
	for i, _ := range slc {
		slc[i] = r.Rand.NormFloat64()
	}
	u.Mul(chol, u)
	return nil
}
