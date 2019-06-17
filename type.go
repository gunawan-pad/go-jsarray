package jsarray

import "reflect"

// ForEachFunc is ForEach callback function declaration
type ForEachFunc func(item interface{}, index int, array []interface{})

// FilterFunc is Filter callback function declaration
type FilterFunc func(item interface{}, index int, array []interface{}) bool

// MapFunc is Map callback function declaration
type MapFunc func(item interface{}, index int, array []interface{}) interface{}

// ReduceFunc is Reduce callback function declaration
type ReduceFunc func(accumulator, currentValue interface{}, currentIndex int, array []interface{}) interface{}

// LessFunc is Less callback function declaration
type LessFunc func(firstEl, secondEl interface{}) bool

// Sorter is a helper for sort method
type Sorter struct {
	array    []interface{}
	lessFunc LessFunc
}

// Array struct
type Array struct {
	_array []interface{}
	_type  reflect.Type

	_sorter Sorter
}
