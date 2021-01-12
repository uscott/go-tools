package rndm

/*
// BvrtNorm is a 2 dimensional normal distribution
type BvrtNorm struct {
	Mu1    float64
	Mu2    float64
	Sigma1 float64
	Sigma2 float64
	Rho    float64
}

// NewBvrtNorm returns a pointer to a BvrtNorm
func NewBvrtNorm(mu1, mu2, sig1, sig2, rho float64) *BvrtNorm {
	var bvn BvrtNorm = BvrtNorm{
		Mu1:    mu1,
		Mu2:    mu2,
		Sigma1: sig1,
		Sigma2: sig2,
		Rho:    rho,
	}
	return &bvn
}

// LogDensity is the natural log of the density function
func (bvn *BvrtNorm) LogDensity(x1, x2 float64) float64 {
	m1, m2, s1, s2, c := bvn.Mu1, bvn.Mu2, bvn.Sigma1, bvn.Sigma2, bvn.Rho
	z1, z2 := (x1-m1)/s1, (x2-m2)/s2
	var ld float64 = 0
	ld = -math.Log(2 * math.Pi * s1 * s2 * math.Sqrt(1-c*c))
	ld -= 0.5 / (1 - c*c) * (z1*z1 + z2*z2 - 2*c*z1*z2)
	return ld
}

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
func (src *DfltSrc) Seed(seed int64) {
	rand.Seed(seed)
}

// Uint64 returns random uint64
func (src *DfltSrc) Int63() int64 {
	return rand.Int63()
}

// Seed seeds RNG
func (rng *RNG) Seed(seed int64) {
	rng.Rand.Seed(seed)
}

// Uint64 returns pseudo random uint64
func (rng *RNG) Uint64() uint64 {
	return rng.Rand.Uint64()
}

// NewRngSeeded returns pointer to seeded RNG
func NewRngSeeded(seed int64) *RNG {
	return &RNG{rand.New(rand.NewSource(seed))}
}

// NewRng returns pointer to RNG seeded with clock time
func NewRng() *RNG {
	return NewRngSeeded(uint64(time.Now().UnixNano()))
}

// MvRand stores a multivariate random draw inside eps
func (rng *MvRng) MvRand(chol *mat.TriDense, upper bool, eps []float64) error {
	if chol == nil {
		return errs.NilPtr
	}
	var (
		n int
		u *mat.Dense
	)
	n, _ = chol.Dims()
	if cap(eps) < n {
		return errs.Capacity
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
*/
