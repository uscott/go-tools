package slice

import (
	"fmt"
	"reflect"
	"sort"
)

// FltRev reverses the order of the slice of float64s
func FltRev(slc []float64) {
	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		slc[i], slc[j] = slc[j], slc[i]
	}
}

// Reverse reverses the order of a slice, pointer to slice or pointer to array
// in place (an array argument will be non-addressable)
func Reverse(v interface{}) {

	x := reflect.ValueOf(v)
	switch x.Kind() {
	case reflect.Array:
		panic("Array argument is not addressible - must be pointer to array")
	case reflect.Ptr:
		x = reflect.Indirect(x)
	}

	switch x.Kind() {
	case reflect.Slice:
		swp := reflect.Swapper(x.Interface())
		for i, j := 0, x.Len()-1; i < j; i, j = i+1, j-1 {
			swp(i, j)
		}
	case reflect.Array:
		for i, j := 0, x.Len()-1; i < j; i, j = i+1, j-1 {
			vi, vj := x.Index(i), x.Index(j)
			swp := reflect.ValueOf(vi.Interface())
			vi.Set(vj)
			vj.Set(swp)
		}
	default:
		panic("Argument must be array or slice")
	}

}

func checkSliceArg(q *reflect.Value) {
	if q == nil {
		panic("Nil pointer")
	}

	if q.Kind() != reflect.Ptr {
		panic("First argument must be a pointer to a slice")
	}

	*q = q.Elem()
	if q.Kind() != reflect.Slice {
		panic("First argument must be a pointer to a slice")
	}
}

func rmIndex(q *reflect.Value, j int) {
	switch j {
	case 0:
		*q = q.Slice(1, q.Len())
	case q.Len() - 1:
		*q = q.Slice(0, q.Len()-1)
	default:
		*q = reflect.AppendSlice(q.Slice(0, j), q.Slice(j+1, q.Len()))
	}
}

func RmByIndex(p interface{}, indices ...int) {

	if len(indices) == 0 {
		return
	}

	q := reflect.ValueOf(p)

	checkSliceArg(&q)

	if q.Len() == 0 {
		return
	}

	var ix []int
	if sort.IntsAreSorted(indices) {
		ix = indices
	} else {
		n := len(indices)
		ix = make([]int, n)
		ncpy := copy(ix, indices)
		if ncpy < n {
			panic(fmt.Sprintf("Only %d elements copied instead of %d", ncpy, n))
		}
		sort.Ints(ix)
	}

	if ix[0] < 0 || q.Len() <= ix[len(ix)-1] {
		panic("Index out of bounds")
	}

	nrm := 0
	for _, i := range ix {
		rmIndex(&q, i-nrm)
		nrm++
	}

	reflect.ValueOf(p).Elem().Set(q)
}

// RmByVal removes elements from a slice that are equal to a given value
// p: pointer to slice
// x: value(s) or pointer(s) to value
func RmByVal(p interface{}, x ...interface{}) {

	if len(x) == 0 {
		return
	}

	q := reflect.ValueOf(p)

	checkSliceArg(&q)

	if q.Len() == 0 {
		return
	}

	elemkind := q.Type().Elem().Kind()

	for _, a := range x {

		y := reflect.ValueOf(a)

		if elemkind != reflect.Ptr && y.Kind() == reflect.Ptr {
			y = y.Elem()
		}

		z := y.Interface()
		n, nrm := q.Len(), 0

		for i := 0; i < n; i++ {
			j := i - nrm
			if reflect.DeepEqual(q.Index(j).Interface(), z) {
				rmIndex(&q, j)
				nrm++
			}
		}
	}

	reflect.ValueOf(p).Elem().Set(q)
}

// StrSlcRm removes string s from *s in place
func StrSlcRm(slc *[]string, s string) {
	n, nrm := len(*slc), 0
	for i := 0; i < n; i++ {
		j := i - nrm
		t := (*slc)[j]
		if t == s {
			nrm++
			switch {
			case j == 0:
				*slc = (*slc)[1:]
			case j == len(*slc)-1:
				*slc = (*slc)[:len(*slc)-1]
			default:
				*slc = append((*slc)[:j], (*slc)[j+1:]...)
			}
		}
	}
}

// StrSlcEquals tests whether the two string slices are
// equal by value
func StrSlcEquals(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, w := range s1 {
		if w != s2[i] {
			return false
		}
	}
	return true
}
