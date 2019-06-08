package jsarray

import (
	"reflect"
)

// GetValueKind gets value and kind of interface variable
func GetValueKind(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

// ConvertInterfaceToArrayInterface convert interface{} to []interface{}
func ConvertInterfaceToArrayInterface(arg interface{}) (out []interface{}, ok bool, itemType reflect.Type) {
	slice, success := GetValueKind(arg, reflect.Slice)

	if !success {
		ok = false
		return
	}

	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		// fmt.Println(slice.Index(i).Kind())
		if i == 0 {
			itemType = slice.Index(i).Type() // .Kind()
		}
		out[i] = slice.Index(i).Interface()
	}

	return out, true, itemType
}
