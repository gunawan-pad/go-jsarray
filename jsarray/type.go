package jsarray

import "reflect"

// AUForEachFunc is ForEach callback function declaration
type AUForEachFunc func(item interface{}, index int, array []interface{})

// AUFilterFunc is Filter callback function declaration
type AUFilterFunc func(item interface{}, index int, array []interface{}) bool

// AUMapFunc is Map callback function declaration
type AUMapFunc func(item interface{}, index int, array []interface{}) interface{}

// AUReduceFunc is Reduce callback function declaration
type AUReduceFunc func(accumulator, currentValue interface{}, currentIndex int, array []interface{}) interface{}

// AULessFunc is Less callback function declaration
type AULessFunc func(firstEl, secondEl interface{}) bool

// Sorter is a helper for sort method
type Sorter struct {
	array    []interface{}
	LessFunc AULessFunc
}

// Array struct
type Array struct {
	_array []interface{}
	_type  reflect.Type

	_sorter Sorter
}
