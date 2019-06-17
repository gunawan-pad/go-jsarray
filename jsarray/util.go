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

// IterateNestedArray iterate each element and execute the callback function
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

// IsArray checks if a variable is a slice or an array
func IsArray(variable interface{}) bool {
	kind := reflect.ValueOf(variable).Kind()
	return kind == reflect.Slice || kind == reflect.Array
}

// fixStartEnd fix start and end to valid values
// (positive, >= array length)
func fixStartEnd(start, end, arrLen int) (int, int) {
	if start < 0 {
		if start += arrLen; start < 0 {
			start = 0
		}
	}
	if end < 0 {
		end += arrLen
	}
	if end < 0 {
		end = 0
	}

	if end > arrLen {
		end = arrLen
	}

	// if start > end {
	// 	return 0, pa
	// }

	return start, end
}

func createArrayTest(maxIteration int,
	pfunc func(int, int) (interface{}, []interface{}),
) ([]interface{}, []interface{}) {
	var methodResults, jsarrResults []interface{}

	mi := maxIteration * (-1)
	for i := mi; i < maxIteration; i++ {
		for j := mi; j < maxIteration; j++ {
			methodRes, parrRes := pfunc(i, j) // arr.Splice(start, end)
			methodResults = append(methodResults, methodRes)
			jsarrResults = append(jsarrResults, parrRes)
		}
	}

	return methodResults, jsarrResults
}
