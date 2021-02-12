package maps

import (
	"fmt"
	"reflect"
)

// ClearMap effectively calls delete on all keys
func ClearMap(m interface{}) {

	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic(fmt.Sprintf("Not a map: %+v", m))
	}

	for _, k := range v.MapKeys() {
		v.SetMapIndex(k, reflect.Value{})
	}
}

// ShrinkMapSlices takes a map of slices and shrinks all of them to length 0.
// Panics if m is not a map of slices.
func ShrinkMapSlices(m interface{}) {

	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic(fmt.Sprintf("Not a map: %+v", m))
	}

	if v.Type().Elem().Kind() != reflect.Slice {
		panic("Not a map of slices")
	}

	iter := v.MapRange()
	for iter.Next() {
		if x := iter.Value(); x.Kind() == reflect.Slice {
			v.SetMapIndex(iter.Key(), x.Slice(0, 0))
		}
	}
}

// AllocateMap takes a pointer to a map and allocates the map if it is not already allocated.
// Panics if p is nil or is not a pointer to a map.
func AllocateMap(p interface{}) {

	v := reflect.ValueOf(p)
	kind := v.Kind()
	var m reflect.Value

	if kind != reflect.Ptr || v.IsZero() {
		panic("Non pointer argument, or nil pointer")
	} else {
		m = reflect.Indirect(v)
	}

	if m.Kind() != reflect.Map {
		panic("Not a pointer to a map")
	}

	if !m.IsZero() {
		return
	}

	t := m.Type()
	m.Set(reflect.MakeMap(reflect.MapOf(t.Key(), t.Elem())))
}
