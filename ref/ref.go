package ref

import (
	"reflect"
)

func DeepElem(v reflect.Value) reflect.Value {

	for {
		kind := v.Kind()
		if kind == reflect.Ptr || kind == reflect.Interface {
			v = v.Elem()
		} else {
			break
		}
	}

	return v
}
