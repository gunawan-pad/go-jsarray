package jsarray

import "reflect"

// ForEachFunc is ForEach callback function type
type ForEachFunc func(item interface{}, index int, array []interface{})

// FilterFunc is Filter callback function type
type FilterFunc func(item interface{}, index int, array []interface{}) bool

// MapFunc is Map callback function type
type MapFunc func(item interface{}, index int, array []interface{}) interface{}

// ReduceFunc is Reduce callback function type
type ReduceFunc func(accumulator, currentValue interface{}, currentIndex int, array []interface{}) interface{}

// LessFunc is Less callback function type
type LessFunc func(firstEl, secondEl interface{}) bool

// Array struct
type Array struct {
	_array []interface{}
	_type  reflect.Type

	_sorter Sorter
}

type Itf interface{}
type Aritf []Itf
