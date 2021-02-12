package dec

import (
	"github.com/shopspring/decimal"
	"github.com/uscott/go-tools/errs"
)

// DecimalChunkRound rounds x in place to the nearest chunk
// specified by sz
func DecimalChunkRound(x, sz *decimal.Decimal) (err error) {
	if x == nil || sz == nil {
		return errs.NilPtrArg
	}
	*x = sz.Mul(x.Div(*sz).Round(0))
	return
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
