package jsarray

import (
	"fmt"
)

func Filter(filterFunc AUFilterFunc, array []interface{}) []interface{} {

	var returnArray []interface{}
	for idx, item := range array {
		if filterFunc(item, idx, array) {
			returnArray = append(returnArray, item)
		}
	}

	return returnArray
}

func Find(filterFunc AUFilterFunc, array []interface{}) interface{} {

	for idx, item := range array {
		if filterFunc(item, idx, array) {
			return item
		}
	}

	return nil
}

func FindIndex(filterFunc AUFilterFunc, array []interface{}) interface{} {

	for idx, item := range array {
		if filterFunc(item, idx, array) {
			return idx
		}
	}

	return -1
}

func Map(mapFunc AUMapFunc, array []interface{}) []interface{} {

	var returnArray []interface{}
	for idx, item := range array {
		returnArray = append(returnArray, mapFunc(item, idx, array))
	}

	return returnArray
}

func Reduce(reduceFunc AUReduceFunc, initialValue interface{}, array []interface{}) interface{} {

	var ret = initialValue
	for idx, item := range array {
		ret = reduceFunc(ret, item, idx, array)
	}

	return ret
}

func ReduceRight(reduceFunc AUReduceFunc, initialValue interface{}, array []interface{}) interface{} {

	var ret = initialValue
	for idx := len(array) - 1; idx >= 0; idx-- {
		item := array[idx]
		ret = reduceFunc(ret, item, idx, array)
	}

	return ret
}

func Some(filterFunc AUFilterFunc, array []interface{}) bool {

	for idx, item := range array {
		if filterFunc(item, idx, array) {
			return true
		}
	}

	return false
}

func Every(filterFunc AUFilterFunc, array []interface{}) bool {
	for idx, item := range array {
		if !filterFunc(item, idx, array) {
			return false
		}
	}

	return true
}

func Join(delimiter string, array []interface{}) string {
	ret, del := "", ""

	for index, item := range array {
		if index > 0 {
			del = delimiter
		}
		ret += fmt.Sprintf("%s%v", del, item) // item.(string)
	}

	return ret
}

////////////////////////////////////////////////////////
