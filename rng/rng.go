package rng

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

type RNG struct {
	*rand.Rand
}

func (rng RNG) Seed(seed uint64) {
	rng.Rand.Seed(int64(seed))
}

func (rng RNG) Uint64() uint64 {
	return rng.Rand.Uint64()
}

func NewRng() RNG {
	return RNG{rand.New(rand.NewSource(time.Now().UnixNano()))}
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
