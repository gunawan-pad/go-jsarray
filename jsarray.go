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

// NewArray is constructor, creates a Array from array of any type
func NewArray(array interface{}) *Array {
	// kind: array item data type

	// create an array copy
	arrInterface, ok, _ := ConvertInterfaceToArrayInterface(array)
	if !ok {
		log.Fatal("Error jsarray.NewArray conversion to []interface{}")
	}
	return (*Array)(&arrInterface)
}

// NewArrayFromInterfaceArray is constructor, creates a Array from array of interface
func NewArrayFromInterfaceArray(array []interface{} /*, dontCreateCopy bool*/) *Array {
	dontCreateCopy := false

	var arrayCopy []interface{}
	if !dontCreateCopy {
		// create an array copy
		arrayCopy = make([]interface{}, len(array))
		copy(arrayCopy, array)
	} else {
		// Used as reference
		arrayCopy = array
	}
	return (*Array)(&arrayCopy)
}

// GetResult to get the result/internal array (array of interface{})
func (pa *Array) GetResult() []interface{} {
	return *pa
}

// GetArray to get the result/internal array (array of interface{})
func (pa *Array) GetArray() []interface{} {
	return *pa
}

// Map method creates a new array with the results of calling a provided function
// on every element in the calling array.
func (pa *Array) Map(callbackfn MapFunc) *Array {
	// arrLen := len(*pa)
	ra := *pa // make([]interface{}, arrLen)
	for i, v := range *pa {
		ra[i] = callbackfn(v, i, *pa)
	}

	return pa // (*Array)(&ra)
}

// ForEach method executes a provided function once for each array element.
func (pa *Array) ForEach(callbackfn ForEachFunc) {
	array := *pa

	for idx, item := range array {
		callbackfn(item, idx, array)
	}
}

// Find method returns the value of the first element in the array
// that satisfies the provided testing function. Otherwise nil is returned.
func (pa *Array) Find(predicate FilterFunc) interface{} {
	array := *pa

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
func (pa *Array) FindIndex(predicate FilterFunc) int {
	array := *pa

	for idx, item := range array {
		if predicate(item, idx, array) {
			return idx
		}
	}

	return -1
}

// Filter method creates a new array with all elements that pass the test
// implemented by the provided function
func (pa *Array) Filter(callbackfn FilterFunc) *Array {

	var returnArray []interface{}
	array := *pa

	for idx, item := range array {
		if callbackfn(item, idx, array) {
			returnArray = append(returnArray, item)
		}
	}

	*pa = returnArray
	return pa
}

// Reduce method executes a reducer function (that you provide)
// on each element of the array, resulting in a single output value.
func (pa *Array) Reduce(callbackfn ReduceFunc, initialValue interface{}) interface{} {
	var ret = initialValue
	array := *pa

	for idx, item := range array {
		ret = callbackfn(ret, item, idx, array)
	}

	return ret
}

// ReduceRight method applies a function against an accumulator
// and each value of the array (from right-to-left) to reduce it
// to a single value.
func (pa *Array) ReduceRight(callbackfn ReduceFunc, initialValue interface{}) interface{} {
	var ret = initialValue
	array := *pa

	for idx := len(array) - 1; idx >= 0; idx-- {
		ret = callbackfn(ret, array[idx], idx, array)
	}

	return ret
}

// Some method tests whether at least one element in the array
// passes the test implemented by the provided function. It returns a Boolean value.
// Note: This method returns false for any condition put on an empty array.
func (pa *Array) Some(callbackfn FilterFunc) bool {
	array := *pa

	for idx, item := range array {
		if callbackfn(item, idx, array) {
			return true
		}
	}

	return false
}

// Every method tests whether all elements in the array pass
// the test implemented by the provided function. It returns a Boolean value.
// Note: This method returns true for any condition put on an empty array.
func (pa *Array) Every(callbackfn FilterFunc) bool {
	array := *pa
	for idx, item := range array {
		if !callbackfn(item, idx, array) {
			return false
		}
	}

	return true
}

// Join method creates and returns a new string by concatenating
// all of the elements in an array (or an array-like object), separated by separator string.
// If the array has only one item, then that item will be returned without using the separator.
func (pa *Array) Join(separator string) string {
	array := *pa
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

// Includes method determines whether an array includes a certain value
// among its entries, returning true or false as appropriate.
func (pa *Array) Includes(searchElement interface{}) bool {
	return pa.IndexOf(searchElement, 0) > -1
}

// IndexOf method returns the first index at which a given element
// can be found in the array, or -1 if it is not present.
func (pa *Array) IndexOf(searchElement interface{}, fromIndex int) int {
	array := *pa
	// TODO: check fromIndex
	if fromIndex < 0 {
		if fromIndex += len(array); fromIndex < 0 {
			fromIndex = 0
		}
	}

	for index, item := range array {
		if searchElement == item && index >= fromIndex {
			return index
		}
	}

	return -1
}

// LastIndexOf method returns the last index at which a given element
// can be found in the array, or -1 if it is not present. The array is
// searched backwards, starting at fromIndex.
func (pa *Array) LastIndexOf(searchElement interface{}, fromIndex int) int {
	array := *pa
	arrLen := len(array)

	if fromIndex > (arrLen - 1) {
		fromIndex = arrLen - 1
	}
	if fromIndex < 0 {
		if fromIndex += arrLen; fromIndex < 0 {
			return -1
		}
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
// from a start index (default zero) to an end index (default array length)
// with a static value. It returns the modified array.
func (pa *Array) Fill(value interface{}, start, end int) *Array {
	array := *pa

	// if start < 0 || end > len(array) {
	// 	panic("jsarray Fill start must be >=0 and end must be <= array length")
	// }

	arrLen := len(array)

	start, end = fixStartEnd(start, end, arrLen)

	if start > end {
		return pa
	}

	for index := range array {
		if index >= start && index < end {
			array[index] = value
		}
	}

	// pCopy := make([]interface{}, end-start) //  array[start:end]
	// for i := range pCopy {
	// 	pCopy[i] = value
	// }
	// copy(array[start:], pCopy)

	return pa
}

// Reverse method reverses an array in place. The first array element
// becomes the last, and the last array element becomes the first.
func (pa *Array) Reverse() *Array {
	array := *pa
	var returnArray = make([]interface{}, len(array))

	j := 0
	for idx := len(array) - 1; idx >= 0; idx-- {
		returnArray[j] = array[idx]
		j++
	}

	*pa = returnArray
	return pa
}

// Flat method creates a new array with all sub-array elements concatenated
// into it recursively up to the specified depth.
func (pa *Array) Flat(depth int) *Array {
	res := FlatArray(*pa, depth)
	*pa = res

	return pa
}

// Sort method sorts the elements of an array in place and returns
// the sorted array.
func (pa *Array) Sort(comparefn LessFunc) *Array {
	// sorter := new(Sorter)
	sorter := &Sorter{
		array:    *pa,
		lessFunc: comparefn,
	}

	sort.Sort(sorter)

	return pa
}

// Shift method removes the first element from an array and returns
// that removed element. This method changes the length of the array.
func (pa *Array) Shift() interface{} {
	if len(*pa) == 0 {
		return nil
	}

	ret := (*pa)[0]
	*pa = (*pa)[1:]

	return ret
}

// Unshift method adds one or more elements to the beginning of an array
// and returns the new length of the array.
func (pa *Array) Unshift(elements ...interface{}) int {
	*pa = append(elements, *pa...)

	return len(*pa)
}

// Pop method removes the last element from an array and
// returns that element. This method changes the length of the array.
func (pa *Array) Pop() interface{} {

	if len(*pa) == 0 {
		return nil
	}

	pos := len(*pa) - 1
	ret := (*pa)[pos]
	*pa = (*pa)[:pos]

	return ret
}

// Push method adds one or more elements to the end of an array and
// returns the new length of the array.
func (pa *Array) Push(elements ...interface{}) int {
	*pa = append(*pa, elements...)

	return len(*pa)
}

// Slice method returns a shallow copy of a portion of an array
// into a new array object selected from begin to end (end not included).
// The original array will not be modified.
func (pa *Array) Slice(start, end int) []interface{} {
	arrLen := len(*pa)

	start, end = fixStartEnd(start, end, arrLen)
	// fmt.Println(start, end)

	if start > end {
		return []interface{}{}
	}
	return (*pa)[start:end]
}

// Length method returns the length of the internal array
func (pa *Array) Length() int {
	return len(*pa)
}

// Splice method changes the contents of an array by removing or
// replacing existing elements and/or adding new elements in place.
func (pa *Array) Splice(start, deleteCount int, items ...interface{}) []interface{} {
	arrLen := len(*pa)

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

	resArray := (*pa)[a:b]

	// remainArray := []interface{}{}
	// remainArray = append(remainArray, (*pa)[0:a]...)
	// remainArray = append(remainArray, items...)
	// remainArray = append(remainArray, (*pa)[b:]...)

	// faster
	l1, l2, l3 := len((*pa)[0:a]), len(items), len((*pa)[b:])
	newLen := l1 + l2 + l3
	remainArray := make([]interface{}, newLen)
	copy(remainArray[0:l1], (*pa)[0:a])
	copy(remainArray[l1:l1+l2], items)
	copy(remainArray[l1+l2:newLen], (*pa)[b:])

	*pa = remainArray
	return resArray
}

// Shuffle shuffles (randomize the order of the elements in)
// an array (in place)
func (pa *Array) Shuffle() *Array {
	array := *pa
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
	arrLen := len(*pa)
	array := *pa
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

	*pa = uniqArr[:j:j]
	return pa
}

// Chunk method split an array into chunks (as in PHP array_chunk).
// Size must be greater than 0.
// Returns array of array of interface ([][]interface{}) and error
func (pa *Array) Chunk(size int) ([][]interface{}, error) {
	if size < 1 {
		return nil, errors.New("Array.Chunk size value must be > 0")
	}

	array := *pa
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

// Concat method is used to merge two or more arrays. This method does not
// change the existing arrays, but instead returns a new array.
func (pa *Array) Concat(items ...interface{}) []interface{} {
	resArr := *pa
	resArr = append(resArr, items[0].([]interface{})...)

	return resArr
}

// CopyWithin method shallow copies part of an array to another location
// in the same array and returns it without modifying its length.
//
// The CopyWithin method is a mutable method. It does not alter the length
// of the array, but it will change its content if necessary.
func (pa *Array) CopyWithin(target, start, end int) []interface{} {
	array := *pa
	arrLen := len(array)

	if target < 0 {
		if target += arrLen; target < 0 {
			target = 0
		}
	}
	if target > arrLen {
		return array
	}

	start, end = fixStartEnd(start, end, arrLen)

	if start > end {
		return array
	}

	// fmt.Println(start, start, end)
	pCopy := array[start:end]

	// resArr := make([]interface{}, len(array))
	// copy(resArr, array)
	resArr := array // in place, *pa modified
	copy(resArr[target:], pCopy)

	// p1 := array[0:target]
	// p2 := array[start+len(pCopy)+1:]
	// resArr := []interface{}{}
	// resArr = append(resArr, p1...)
	// resArr = append(resArr, pCopy...)
	// resArr = append(resArr, p2...)

	return resArr[:arrLen]
}

// GetItem method gets item at specific index
func (pa *Array) GetItem(index int) interface{} {
	return (*pa)[index]
}

// SetItem method sets item value at specific index
func (pa *Array) SetItem(index int, value interface{}) {
	// if index < 0 || index > len(*pa) {
	// 	return errors.New("index out of range")
	// }
	(*pa)[index] = value
	// return nil
}
