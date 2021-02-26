package dec

import (
	"github.com/shopspring/decimal"
	"github.com/uscott/go-tools/errs"
)

// ChunkCeil returns *x ceiled to the chunk *sz
// Panics if nil pointer is passed
func ChunkCeil(x, sz *decimal.Decimal) decimal.Decimal {

	if x == nil || sz == nil {
		panic(errs.NilPtrArg)
	}

	return sz.Mul(x.Div(*sz).Ceil())
}

// ChunkFlorr returns *x floored to the chunk *sz
// Panics if nil pointer is passed
func ChunkFloor(x, sz *decimal.Decimal) decimal.Decimal {

	if x == nil || sz == nil {
		panic(errs.NilPtrArg)
	}

	return sz.Mul(x.Div(*sz).Floor())
}

// ChunkRound returns *x rounded to the nearest chunk *sz
// Panics if nil pointer is passed
func ChunkRound(x, sz *decimal.Decimal) decimal.Decimal {

	if x == nil || sz == nil {
		panic(errs.NilPtrArg)
	}

	return sz.Mul(x.Div(*sz).Round(0))
}

func PtrDecimalMax(args ...*decimal.Decimal) (ptrmax *decimal.Decimal) {
	if len(args) == 0 {
		return
	}
	ptrmax = args[0]
	for _, p := range args[1:] {
		if p == nil {
			continue
		}
		if ptrmax == nil {
			ptrmax = p
		}
		if ptrmax.LessThan(*p) {
			ptrmax = p
		}
	}
	return
}

func PtrDecimalMin(args ...*decimal.Decimal) (ptrmin *decimal.Decimal) {
	if len(args) == 0 {
		return
	}
	ptrmin = args[0]
	for _, p := range args[1:] {
		if p == nil {
			continue
		}
		if ptrmin == nil {
			ptrmin = p
		}
		if ptrmin.GreaterThan(*p) {
			ptrmin = p
		}
	}
	return
}
