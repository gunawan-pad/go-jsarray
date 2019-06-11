package jsarray

import "reflect"

// MakeRange creates array of int from start to length -1
func MakeRange(start, length int) []int {
	ra := make([]int, length)
	for i := range ra {
		ra[i] = start + i
	}

	return ra
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

// FlatArray function creates a new array with all sub-array elements concatenated into it recursively up to the specified depth.
func FlatArray(inputArray []interface{}, depth int) (ret []interface{}) {
	ret = []interface{}{}

	IterateNestedArray(inputArray, func(item interface{}) {
		ret = append(ret, item.(interface{}))
	}, 0, depth)

	return
}
