package rndm

import (
	"time"

	"github.com/uscott/gotools/errs"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distuv"
)

// DfltSrc is a wrapper struct for the rand default source
type DfltSrc struct {
}

// RNG is a wrapper struct for rand.Rand
type RNG struct {
	*rand.Rand
}

// MvRng is a source for multivariate random number generation
type MvRng struct {
	Normal *distuv.Normal
	Src    rand.Source
}

// NewMvRng returns a pointer to a seeded MvRng
func NewMvRng(seed uint64) *MvRng {
	var (
		src = &DfltSrc{}
		n   = distuv.Normal{Mu: 0, Sigma: 1, Src: src}
	)
	src.Seed(seed)
	return &MvRng{Normal: &n, Src: src}
}

// Seed seeds MvRng
func (rng *MvRng) Seed(s uint64) {
	rng.Src.Seed(s)
}

// Uint64 returns random uint64
func (rng *MvRng) Uint64() uint64 {
	return rng.Src.Uint64()
}

// Seed seeds DftlSrc
func (src *DfltSrc) Seed(seed uint64) {
	rand.Seed(uint64(seed))
}

// Uint64 returns random uint64
func (src *DfltSrc) Uint64() uint64 {
	return rand.Uint64()
}

// Seed seeds RNG
func (rng *RNG) Seed(seed uint64) {
	rng.Rand.Seed(uint64(seed))
}

// Uint64 returns pseudo random uint64
func (rng *RNG) Uint64() uint64 {
	return rng.Rand.Uint64()
}

// NewRngSeeded returns pointer to seeded RNG
func NewRngSeeded(seed uint64) *RNG {
	return &RNG{rand.New(rand.NewSource(seed))}
}

// NewRng returns pointer to RNG seeded with clock time
func NewRng() *RNG {
	return NewRngSeeded(uint64(time.Now().UnixNano()))
}

// MvRand stores a multivariate random draw inside eps
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
		return errs.ErrCapacity
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
