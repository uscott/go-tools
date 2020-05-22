package rndm

import (
	"github.com/uscott/gotools/misc"
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

type RNG struct {
	*rand.Rand
	Source *rand.Source
}

func (rng RNG) Seed(seed uint64) {
	(*rng.Source).Seed(int64(seed))
}

func (rng RNG) Uint64() uint64 {
	return rng.Rand.Uint64()
}

func NewRngSeeded(seed int) RNG {
	src := rand.NewSource(int64(seed) + time.Now().UnixNano())
 	return RNG{rand.New(src), &src}
}

func NewRng() RNG {
	return NewRngSeeded(0)
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
