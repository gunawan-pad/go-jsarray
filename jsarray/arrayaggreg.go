package jsarray

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
)

// type T interface{}

type ItemArr []interface{}

type AUFilterFunc func(item interface{}, index int, array []interface{}) bool
type AUMapFunc func(item interface{}, index int, array []interface{}) interface{}
type AUReduceFunc func(prev, current interface{}, index int, array []interface{}) interface{}

type AULessFunc func(i, j interface{}) bool

type Sorter struct {
	array    []interface{}
	LessFunc AULessFunc
}

// Len is part of sort.Interface.
func (s *Sorter) Len() int {
	return len(s.array)
}

// Swap is part of sort.Interface.
func (s *Sorter) Swap(i, j int) {
	s.array[i], s.array[j] = s.array[j], s.array[i]
}

// Less is part of sort.Interface.
func (s *Sorter) Less(i, j int) bool {
	return s.LessFunc(s.array[i], s.array[j])
}

type Array struct {
	_array []interface{}
	_type  reflect.Type

	_sorter Sorter
}

// NewArray is constructor
func NewArray(array interface{} /*, arrayLen int*/) *Array {
	// kind: array item data type
	arrInterface, ok, reflType := ConvertInterfaceToArrayInterface(array)
	if !ok {
		log.Fatal("Error arrayaggreg.NewArray conversion to []interface{}")
	}

	return &Array{
		_array: arrInterface,
		_type:  reflType,
	}
}

// NewArrayFromInterfaceArray is constructor from Array of interface
// TODO: create reference
func NewArrayFromInterfaceArray(array []interface{} /*, arrayLen int*/) *Array {
	return &Array{
		_array: array,
		// _array: arrInterface,
		_type: nil,
	}
}

// Get to get the internal array (_array)
// create a copy
func (pa *Array) Get(createCopy bool) interface{} {
	if !createCopy {
		return interface{}(pa._array)
	}

	var res = make([]interface{}, len(pa._array))
	copy(res, pa._array)

	return interface{}(res)

	/*
		slice := reflect.MakeSlice(reflect.SliceOf(array._type), len(res), len(res))
		for i := 0; i < len(res); i++ {
			// v := reflect.Indirect(reflect.ValueOf(res[i]).Convert(array._type))
			// slice = reflect.Append(slice, v)

			// slice.Index(i).Set(reflect.ValueOf(res[i]).Convert(array._type)) //.(array._type)))
			slice.Index(i).Set(reflect.ValueOf(res[i].(array._type)))
		}
		return interface{}(slice)
	*/
}

func (pa *Array) Map(callbackfn AUMapFunc) *Array {
	_array := pa._array
	var returnArray = make([]interface{}, len(_array))

	for idx, item := range _array {
		returnArray[idx] = callbackfn(item, idx, _array)
	}

	pa._array = returnArray
	return pa
}

func (pa *Array) ForEach(callbackfn AUMapFunc) {
	_array := pa._array

	for idx, item := range _array {
		callbackfn(item, idx, _array)
	}
}

func (pa *Array) Find(predicate AUFilterFunc) interface{} {
	_array := pa._array

	for idx, item := range _array {
		if predicate(item, idx, _array) {
			return item
		}
	}

	return nil
}

func (pa *Array) FindIndex(predicate AUFilterFunc) int {
	_array := pa._array

	for idx, item := range _array {
		if predicate(item, idx, _array) {
			return idx
		}
	}

	return -1
}

func (pa *Array) Filter(callbackfn AUFilterFunc) *Array {

	var returnArray []interface{}
	_array := pa._array

	for idx, item := range _array {
		if callbackfn(item, idx, _array) {
			returnArray = append(returnArray, item)
		}
	}

	pa._array = returnArray
	return pa
}

func (pa *Array) Reduce(callbackfn AUReduceFunc, initialValue interface{}) interface{} {
	var ret = initialValue
	_array := pa._array

	for idx, item := range _array {
		ret = callbackfn(ret, item, idx, _array)
	}

	return ret
}

func (pa *Array) ReduceRight(callbackfn AUReduceFunc, initialValue interface{}) interface{} {
	var ret = initialValue
	_array := pa._array

	for idx := len(_array) - 1; idx >= 0; idx-- {
		item := _array[idx]
		ret = callbackfn(ret, item, idx, _array)
	}

	return ret
}

func (pa *Array) Some(callbackfn AUFilterFunc) bool {
	array := pa._array

	for idx, item := range array {
		if callbackfn(item, idx, array) {
			return true
		}
	}

	return false
}

func (pa *Array) Every(callbackfn AUFilterFunc) bool {
	array := pa._array
	for idx, item := range array {
		if !callbackfn(item, idx, array) {
			return false
		}
	}

	return true
}

func (pa *Array) Join(separator string) string {
	array := pa._array
	// var ret bytes.Buffer
	var ret strings.Builder
	del := ""

	for index, item := range array {
		if index > 0 {
			del = separator
		}
		ret.WriteString(fmt.Sprintf("%s%v", del, item))
	}

	return ret.String()
}

func (pa *Array) Includes(searchElement interface{}) bool {
	return pa.IndexOf(searchElement, 0) > -1
}

func (pa *Array) IndexOf(searchElement interface{}, fromIndex int) int {
	array := pa._array

	if fromIndex < 0 {
		fromIndex = len(array) + fromIndex
	}

	for index, item := range array {
		if searchElement == item && index >= fromIndex {
			return index
		}
	}

	return -1
}

func (pa *Array) LastIndexOf(searchElement interface{}, fromIndex int) int {
	array := pa._array
	if fromIndex > (len(array) - 1) {
		fromIndex = len(array) - 1
	}
	if fromIndex < 0 {
		fromIndex = len(array) + fromIndex
	}

	// fmt.Printf("fi:%d\n", fromIndex)

	for index := fromIndex; index >= 0; index-- {
		if searchElement == array[index] {
			return index
		}
	}

	return -1
}

func (pa *Array) Fill(value interface{}, start, end int) *Array {
	array := pa._array
	for index := range array {
		if index >= start && index <= end {
			array[index] = value
		}
	}

	return pa
}

func (pa *Array) Reverse() *Array {
	array := pa._array
	var returnArray = make([]interface{}, len(array))

	j := 0
	for idx := len(array) - 1; idx >= 0; idx-- {
		returnArray[j] = array[idx]
		j++
	}

	pa._array = returnArray
	return pa
}

func IterateNestedArray(
	arrIntf []interface{},
	actionFunc func(interface{}),
	currentDepth, maxDepth int,
) {
	for _, item := range arrIntf {
		// fmt.Printf("> %v >%s\n", item, reflect.TypeOf(item).String())

		// check if its an array
		// typeStr := reflect.TypeOf(item).String()
		// if strings.HasPrefix(typeStr, "[]") {
		kind := reflect.ValueOf(item).Kind()
		if (kind == reflect.Slice || kind == reflect.Array) &&
			(currentDepth < maxDepth || maxDepth < 0) {
			IterateNestedArray(item.([]interface{}), actionFunc, currentDepth+1, maxDepth)
		} else {
			actionFunc(item)
		}
	}
}

// FlatArray creates a new array with all sub-array elements
// concatenated into it recursively up to specified depth
func FlatArray(inputArray []interface{}, depth int) (ret []interface{}) {
	ret = []interface{}{}

	IterateNestedArray(inputArray, func(item interface{}) {
		ret = append(ret, item.(interface{}))
	}, 0, depth)

	return
}

// Flat creates a new array with all sub-array elements
// concatenated into it recursively up to specified depth
func (pa *Array) Flat(depth int) *Array {
	res := FlatArray(pa._array, depth)
	pa._array = res

	return pa
}

func (pa *Array) Sort(comparefn AULessFunc) *Array {
	array := pa._array

	pa._sorter.array = array
	pa._sorter.LessFunc = comparefn

	sort.Sort(&pa._sorter)

	return pa
}

// TODO:
// -concat, flatMap , sort
// lastIndexOf, pop,push,
// shift,slice, splice
