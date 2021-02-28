package dec

import (
	"github.com/shopspring/decimal"
)

// ChunkCeil returns x ceiled to the chunk sz
// Panics if sz is zero
func ChunkCeil(x, sz decimal.Decimal) decimal.Decimal {
	if sz.IsZero() {
		panic("Division by zero")
	}
	return sz.Mul(x.Div(sz).Ceil())
}

// ChunkFloor returns x floored to the chunk sz
// Panics if sz is zero
func ChunkFloor(x, sz decimal.Decimal) decimal.Decimal {
	if sz.IsZero() {
		panic("Division by zero")
	}
	return sz.Mul(x.Div(sz).Floor())
}

// ChunkRound returns x rounded to the nearest chunk sz
// Panics if sz is zero
func ChunkRound(x, sz decimal.Decimal) decimal.Decimal {
	if sz.IsZero() {
		panic("Division by zero")
	}
	return sz.Mul(x.Div(sz).Round(0))
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
