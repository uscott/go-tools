package rndm

import (
	"fmt"
	"time"

	"github.com/uscott/gotools/errs"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distuv"
)

type DfltSrc struct {
}

type RNG struct {
	*rand.Rand
}

type MvRng struct {
	Normal *distuv.Normal
	Src    rand.Source
}

func NewMvRng(seed uint64) *MvRng {
	var (
		src = rand.NewSource(seed)
		n   = distuv.Normal{Mu: 0, Sigma: 1, Src: src}
	)
	return &MvRng{Normal: &n, Src: src}
}

func (rng *MvRng) Seed(s uint64) {
	rng.Src.Seed(s)
}

func (rng *MvRng) Uint64() uint64 {
	return rng.Src.Uint64()
}

func (src *DfltSrc) Seed(seed uint64) {
	rand.Seed(uint64(seed))
}

func (src *DfltSrc) Uint64() uint64 {
	return rand.Uint64()
}

func (rng *RNG) Seed(seed uint64) {
	rng.Rand.Seed(uint64(seed))
}

func (rng *RNG) Uint64() uint64 {
	return rng.Rand.Uint64()
}

func NewRngSeeded(seed uint64) *RNG {
	return &RNG{rand.New(rand.NewSource(seed))}
}

func NewRng() *RNG {
	return NewRngSeeded(uint64(time.Now().UnixNano()))
}

func (rng *MvRng) MvRand(chol *mat.TriDense, upper bool, eps []float64) error {
	if chol == nil {
		return errs.ErrNilPtr
	}
	var (
		n int
		u *mat.Dense
	)
	n, _ = chol.Dims()
	if cap(eps) < n {
		return fmt.Errorf("Buffer has insufficient capacity")
	}
	if upper {
		u = mat.NewDense(1, n, eps)
	} else {
		u = mat.NewDense(n, 1, eps)
	}
	for i := range eps {
		eps[i] = rng.Normal.Rand()
	}
	if upper {
		u.Mul(u, chol)
	} else {
		u.Mul(chol, u)
	}
	return nil
}
