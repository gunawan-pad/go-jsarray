package jsarray

import "reflect"

type AUForEachFunc func(item interface{}, index int, array []interface{})
type AUFilterFunc func(item interface{}, index int, array []interface{}) bool
type AUMapFunc func(item interface{}, index int, array []interface{}) interface{}
type AUReduceFunc func(accumulator, currentValue interface{}, currentIndex int, array []interface{}) interface{}

type AULessFunc func(firstEl, secondEl interface{}) bool

// Sorter is a helper for sort method
type Sorter struct {
	array    []interface{}
	LessFunc AULessFunc
}

type Array struct {
	_array []interface{}
	_type  reflect.Type

	_sorter Sorter
}
