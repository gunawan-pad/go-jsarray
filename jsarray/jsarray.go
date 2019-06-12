package jsarray

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// NewArray is constructor, creates Array from array of any type
func NewArray(array interface{} /*, arrayLen int*/) *Array {
	// kind: array item data type
	arrInterface, ok, reflType := ConvertInterfaceToArrayInterface(array)
	if !ok {
		log.Fatal("Error jsarray.NewArray conversion to []interface{}")
	}

	return &Array{
		_array: arrInterface,
		_type:  reflType,
	}
}

// NewArrayFromInterfaceArray is constructor, creates Array from array of interface
func NewArrayFromInterfaceArray(array []interface{} /*, arrayLen int*/) *Array {
	return &Array{
		_array: array,
		// _array: arrInterface,
		_type: nil,
	}
}

// Get to get the result/internal array as interface{}.
// Obsolete method, use GetResult instead
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

// GetResult to get the result/internal array (array of interface{})
func (pa *Array) GetResult() []interface{} {
	return pa._array
}

// Map method creates a new array with the results of calling a provided function
// on every element in the calling array.
func (pa *Array) Map(callbackfn AUMapFunc) *Array {
	array := pa._array
	var returnArray = make([]interface{}, len(array))

	for idx, item := range array {
		returnArray[idx] = callbackfn(item, idx, array)
	}

	pa._array = returnArray
	return pa
}

// ForEach method executes a provided function once for each array element.
func (pa *Array) ForEach(callbackfn AUForEachFunc) {
	array := pa._array

	for idx, item := range array {
		callbackfn(item, idx, array)
	}
}

// Find method returns the value of the first element in the array
// that satisfies the provided testing function. Otherwise nil is returned.
func (pa *Array) Find(predicate AUFilterFunc) interface{} {
	array := pa._array

	for idx, item := range array {
		if predicate(item, idx, array) {
			return item
		}
	}

	return nil
}

// FindIndex method returns the index of the first element in the array
// that satisfies the provided testing function. Otherwise, it returns -1,
// indicating that no element passed the test.
func (pa *Array) FindIndex(predicate AUFilterFunc) int {
	array := pa._array

	for idx, item := range array {
		if predicate(item, idx, array) {
			return idx
		}
	}

	return -1
}

// Filter method creates a new array with all elements that pass the test
// implemented by the provided function
func (pa *Array) Filter(callbackfn AUFilterFunc) *Array {

	var returnArray []interface{}
	array := pa._array

	for idx, item := range array {
		if callbackfn(item, idx, array) {
			returnArray = append(returnArray, item)
		}
	}

	pa._array = returnArray
	return pa
}

// Reduce method executes a reducer function (that you provide) on each element of the array, resulting in a single output value.
func (pa *Array) Reduce(callbackfn AUReduceFunc, initialValue interface{}) interface{} {
	var ret = initialValue
	array := pa._array

	for idx, item := range array {
		ret = callbackfn(ret, item, idx, array)
	}

	return ret
}

// ReduceRight method applies a function against an accumulator and each value of the array (from right-to-left) to reduce it to a single value.
func (pa *Array) ReduceRight(callbackfn AUReduceFunc, initialValue interface{}) interface{} {
	var ret = initialValue
	array := pa._array

	for idx := len(array) - 1; idx >= 0; idx-- {
		item := array[idx]
		ret = callbackfn(ret, item, idx, array)
	}

	return ret
}

// Some method tests whether at least one element in the array passes the test implemented by the provided function. It returns a Boolean value.
// Note: This method returns false for any condition put on an empty array.
func (pa *Array) Some(callbackfn AUFilterFunc) bool {
	array := pa._array

	for idx, item := range array {
		if callbackfn(item, idx, array) {
			return true
		}
	}

	return false
}

// Every method tests whether all elements in the array pass the test implemented by the provided function. It returns a Boolean value.
// Note: This method returns true for any condition put on an empty array.
func (pa *Array) Every(callbackfn AUFilterFunc) bool {
	array := pa._array
	for idx, item := range array {
		if !callbackfn(item, idx, array) {
			return false
		}
	}

	return true
}

// Join method creates and returns a new string by concatenating all of the elements in an array (or an array-like object), separated by separator string.
// If the array has only one item, then that item will be returned without using the separator.
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

// Includes method determines whether an array includes a certain value among its entries, returning true or false as appropriate.
func (pa *Array) Includes(searchElement interface{}) bool {
	return pa.IndexOf(searchElement, 0) > -1
}

// IndexOf method returns the first index at which a given element
// can be found in the array, or -1 if it is not present.
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

// LastIndexOf method returns the last index at which a given element
// can be found in the array, or -1 if it is not present. The array is searched backwards, starting at fromIndex.
func (pa *Array) LastIndexOf(searchElement interface{}, fromIndex int) int {
	array := pa._array
	arrLen := len(array)

	if fromIndex > (arrLen - 1) {
		fromIndex = arrLen - 1
	}
	if fromIndex < 0 {
		fromIndex = arrLen + fromIndex
	}

	// fmt.Printf("fi:%d\n", fromIndex)

	for index := fromIndex; index >= 0; index-- {
		if searchElement == array[index] {
			return index
		}
	}

	return -1
}

// Fill method fills (modifies) all the elements of an array
// from a start index (default zero) to an end index (default array length) with a static value. It returns the modified array.
func (pa *Array) Fill(value interface{}, start, end int) *Array {
	array := pa._array
	for index := range array {
		if index >= start && index <= end {
			array[index] = value
		}
	}

	return pa
}

// Reverse method reverses an array in place. The first array element
// becomes the last, and the last array element becomes the first.
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

// Flat method creates a new array with all sub-array elements concatenated into it recursively up to the specified depth.
func (pa *Array) Flat(depth int) *Array {
	res := FlatArray(pa._array, depth)
	pa._array = res

	return pa
}

// Sort method sorts the elements of an array in place and returns
// the sorted array.
func (pa *Array) Sort(comparefn AULessFunc) *Array {
	array := pa._array

	pa._sorter.array = array
	pa._sorter.LessFunc = comparefn

	sort.Sort(&pa._sorter)

	return pa
}

// Shift method removes the first element from an array and returns
// that removed element. This method changes the length of the array.
func (pa *Array) Shift() interface{} {
	if len(pa._array) == 0 {
		return nil
	}

	ret := pa._array[0]
	pa._array = pa._array[1:]

	return ret
}

// Unshift method adds one or more elements to the beginning of an array
// and returns the new length of the array.
func (pa *Array) Unshift(elements ...interface{}) int {
	pa._array = append(elements, pa._array...)

	return len(pa._array)
}

// Pop method removes the last element from an array and
// returns that element. This method changes the length of the array.
func (pa *Array) Pop() interface{} {

	if len(pa._array) == 0 {
		return nil
	}

	pos := len(pa._array) - 1
	ret := pa._array[pos]
	pa._array = pa._array[:pos]

	return ret
}

// Push method adds one or more elements to the end of an array and
// returns the new length of the array.
func (pa *Array) Push(elements ...interface{}) int {
	pa._array = append(pa._array, elements...)

	return len(pa._array)
}

// Slice method returns a shallow copy of a portion of an array
// into a new array object selected from begin to end (end not included).
// The original array will not be modified.
func (pa *Array) Slice(begin, end int) []interface{} {
	arrLen := len(pa._array)

	if begin < 0 {
		if begin += arrLen; begin < 0 {
			begin = 0
		}
	}

	if end < 0 {
		end += arrLen
	}

	if end > arrLen {
		end = arrLen
	}

	// fmt.Println(begin, end)

	if begin > end {
		return []interface{}{}
	}
	return pa._array[begin:end]
}

// Length method returns the length of the internal array
func (pa *Array) Length() int {
	return len(pa._array)
}

func (pa *Array) Splice(start, deleteCount int, items ...interface{}) []interface{} {
	arrLen := len(pa._array)

	if start < 0 {
		if start += arrLen; start < 0 {
			start = 0
		}
	}

	if start > arrLen {
		start = arrLen
	}

	if deleteCount < 0 {
		deleteCount = 0
	}

	if deleteCount > arrLen {
		deleteCount = arrLen
	}

	// fmt.Println(start, deleteCount)

	a := start
	b := start + deleteCount

	if b > arrLen {
		b = arrLen
	}

	if b < 0 {
		b = 0
	}

	resArray := pa._array[a:b]

	// remainArray := []interface{}{}
	// remainArray = append(remainArray, pa._array[0:a]...)
	// remainArray = append(remainArray, items...)
	// remainArray = append(remainArray, pa._array[b:]...)

	// faster
	l1, l2, l3 := len(pa._array[0:a]), len(items), len(pa._array[b:])
	newLen := l1 + l2 + l3
	remainArray := make([]interface{}, newLen)
	copy(remainArray[0:l1], pa._array[0:a])
	copy(remainArray[l1:l1+l2], items)
	copy(remainArray[l1+l2:newLen], pa._array[b:])

	pa._array = remainArray
	return resArray
}

// Shuffle shuffles (randomize the order of the elements
// in) an array (in place)
func (pa *Array) Shuffle() *Array {
	array := pa._array
	arrLen := len(array)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i, rnd := range r.Perm(arrLen) {
		t := array[i]
		array[i] = array[rnd]
		array[rnd] = t
	}

	return pa
}

// Unique method remove duplicate values from an array.
func (pa *Array) Unique() *Array {
	arrLen := len(pa._array)
	array := pa._array
	uniqArr := make([]interface{}, arrLen)
	set := make(map[interface{}]bool)

	j := 0
	for _, v := range array {
		if !set[v] {
			set[v] = true
			uniqArr[j] = v
			j++
		}
	}

	pa._array = uniqArr[:j:j]
	return pa
}

// Chunk method split an array into chunks (as in PHP array_chunk).
// Size must be greater than 0.
// Returns array of array of interface ([][]interface{}) and error
func (pa *Array) Chunk(size int) ([][]interface{}, error) {
	if size < 1 {
		return nil, errors.New("Array.Chunk size value must be > 0")
	}

	array := pa._array
	arrLen := len(array)
	chunkedLen := (arrLen / size) + (arrLen % size)
	chunked := make([][]interface{}, chunkedLen)

	j := 0
	for index := 0; index < arrLen; index += size {
		i2 := index + size
		if i2 > arrLen {
			i2 = arrLen
		}
		// chunked = append(chunked, array[index:i2])
		chunked[j] = array[index:i2]
		j++
	}

	return chunked, nil
}

// Split method split an array into chunks, same as Chunk method. See Chunk
// documentation.
func (pa *Array) Split(size int) ([][]interface{}, error) {
	return pa.Chunk(size)
}

func (pa *Array) Concat(items ...interface{}) []interface{} {
	resArr := pa._array
	resArr = append(resArr, items[0].([]interface{})...)

	// for _, item := range items {
	// 	// kind := reflect.ValueOf(item).Kind()
	// 	// if kind == reflect.Slice || kind == reflect.Array {
	// 	// 	// appendItems(item.([]interface{}))
	// 	// 	resArr = append(resArr, item.([]interface{}))
	// 	// } else {
	// 	resArr = append(resArr, item)
	// 	// }
	// 	// 	fmt.Println("it:", item)
	// }

	return resArr
}

func (pa *Array) CopyWithin(target, start, end int) []interface{} {
	array := pa._array
	arrLen := len(array)

	if target < 0 {
		if target += arrLen; target < 0 {
			target = 0
		}
	}
	if target > arrLen {
		return array
	}

	if start < 0 {
		if start += arrLen; start < 0 {
			start = 0
		}
	}

	if end < 0 {
		end += arrLen
	}

	if end > arrLen {
		end = arrLen
	}

	if start > end {
		return array
	}

	// fmt.Println(start, start, end)
	pCopy := array[start:end]
	p1 := array[0:target]
	p2 := array[start+len(pCopy)+1:]

	resArr := []interface{}{}
	resArr = append(resArr, p1...)
	resArr = append(resArr, pCopy...)
	resArr = append(resArr, p2...)

	return resArr[:arrLen]
}
