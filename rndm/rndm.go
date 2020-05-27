package rndm

import (
	"math/rand"
	"time"

	"github.com/uscott/gotools/errs"
	"gonum.org/v1/gonum/mat"
)

type DfltSrc struct {
}

type RNG struct {
	*rand.Rand
}

func (src DfltSrc) Seed(seed uint64) {
	rand.Seed(int64(seed))
}

func (src DfltSrc) Uint64() uint64 {
	return rand.Uint64()
}

func (rng RNG) Seed(seed uint64) {
	rng.Rand.Seed(int64(seed))
}

func (rng RNG) Uint64() uint64 {
	return rng.Rand.Uint64()
}

func NewRngSeeded(seed int) RNG {
	src := rand.NewSource(int64(seed) + time.Now().UnixNano())
	return RNG{rand.New(src)}
}

func NormRand(chol *mat.TriDense, eps []float64) error {
	if chol == nil || eps == nil {
		return errs.ErrNilPtr
	}
	n, _ := chol.Dims()
	u := mat.NewDense(n, 1, eps)
	slc := u.RawMatrix().Data
	for i := range slc {
		slc[i] = rand.NormFloat64()
	}
	u.Mul(chol.T(), u)
	return nil
}
