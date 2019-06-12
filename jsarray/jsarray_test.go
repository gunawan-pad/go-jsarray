package jsarray

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	array1 = []int{1, 2, 3, 4, 5, 4, 6}         // int array
	array2 = []interface{}{1, 2, 3, 4, 5, 4, 6} // interface array

	arrString = []string{
		"satu", "dua", "tiga", "empat", "lima", "empat", "enam",
	}

	arrayNested = []interface{}{
		12, 23,
		[]interface{}{31, 32, 33, []interface{}{331, 332}, 34, 35},
		41, 41,
		[]interface{}{51, 52, 53},
	}

	arrayBig = []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
		60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
		80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
		100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119,
		120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
		140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159,
		160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179,
		180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199,
	}
)

func CompareArray(arrResult, arrCompare []interface{}) bool {
	return !reflect.DeepEqual(arrResult, arrCompare)

	for i, _ := range arrResult {
		if arrResult[i] != arrCompare[i] {
			return true
		}
	}
	return false
}

func TestAGIndexOf(t *testing.T) {
	// arr := NewArrayFromInterfaceArray(array2)
	arr := NewArray(arrString)

	arrResult := []interface{}{}
	for i := -10; i < 10; i++ {
		f := arr.IndexOf("empat", i)
		arrResult = append(arrResult, f)
	}

	fmt.Println(arrResult)
	var arrCompare = []interface{}{3, 3, 3, 3, 3, 3, 3, 5, 5, -1, 3,
		3, 3, 3, 5, 5, -1, -1, -1, -1}
	err := CompareArray(arrResult, arrCompare)

	if err {
		// _, file, line, _ := runtime.Caller(1)
		// fname := filepath.Base(file)
		// t.Fatalf("failed assertion at %s:%d: %s (no error)\n", fname, line, s)
		// t.Fatalf("Wrong number of breakpoints returned for location <%s> (got %d, expected %d)", loc, len(locs), count)

		t.Errorf("Test fails, %s", "AGIndexOf")
	}

}

func TestAGLastIndexOf(t *testing.T) {
	arr := NewArrayFromInterfaceArray(array2)
	// arrRes := arr.LastIndexOf(4, -1)

	arrResult := []interface{}{}
	for i := -10; i < 10; i++ {
		f := arr.LastIndexOf(4, i)
		// fmt.Println(i, f)
		arrResult = append(arrResult, f)
	}

	fmt.Println(arrResult)
	err := false
	var arrCompare = []interface{}{-1, -1, -1, -1, -1, -1, 3, 3, 5, 5, -1, -1, -1, 3, 3, 5, 5, 5, 5, 5}

	err = CompareArray(arrResult, arrCompare)

	if err {
		t.Errorf("Test fails, %s", "AGLastIndexOf")
	}

}

func TestAGMap(t *testing.T) {
	callbackfn := func(item interface{}, index int, array []interface{}) interface{} {
		return item.(int) * 2
	}

	arr := NewArrayFromInterfaceArray(array2)
	arrResult := arr.
		Map(callbackfn).
		GetResult()

	fmt.Println(arrResult)
	var arrCompare = []interface{}{2, 4, 6, 8, 10, 8, 12}
	err := CompareArray(arrResult, arrCompare)

	if err {
		t.Errorf("Test fails, %s", "TestAGMap")
	}

}

func TestAGReduceRight(t *testing.T) {
	arr := NewArrayFromInterfaceArray(array2)
	arrResult := arr.ReduceRight(
		func(tot, item interface{}, index int, array []interface{}) interface{} {
			ii := item.(int) * 2
			return tot.(string) + fmt.Sprintf("%d", ii)
		}, "")

	fmt.Println(arrResult)
	err := arrResult.(string) != "128108642" // "241620161284"
	// "246810812" reduce

	if err {
		t.Errorf("Test fails, %s", "TestAGReduceRight")
	}
}

func TestAGFilter(t *testing.T) {
	arr := NewArray(array1)
	arrResult := arr.
		// Map(func(item interface{}, index int, array []interface{}) interface{} {
		// 	return item.(int) * 2
		// }).
		Filter(func(item interface{}, index int, array []interface{}) bool {
			ii := item.(int)
			return ii > 4
		}).
		GetResult()

	fmt.Println(arrResult)
	var arrCompare = []interface{}{5, 6}
	err := CompareArray(arrResult, arrCompare)

	if err {
		t.Errorf("Test fails, %s", "TestAGFilter")
	}
}

func TestAGJoin(t *testing.T) {
	arr := NewArrayFromInterfaceArray(array2)
	arrResult := arr.Join(", ")

	fmt.Println(arrResult)
	err := arrResult != "1, 2, 3, 4, 5, 4, 6"

	if err {
		t.Errorf("Test fails, %s", "TestAGFilter")
	}
}

func TestAGSort(t *testing.T) {
	// arr := NewArrayFromInterfaceArray(array2)
	arr := NewArray(arrString)
	arrResult := arr.Sort(func(a, b interface{}) bool {
		return a.(string) < b.(string)
	}).GetResult()

	fmt.Println(arrResult)
	var arrCompare = []interface{}{"dua", "empat", "empat", "enam", "lima", "satu", "tiga"}
	// 1, 2, 3, 4, 4, 5, 6}
	err := CompareArray(arrResult, arrCompare)

	if err {
		t.Errorf("Test fails, %s", "TestAGSort")
	}

}

func TestAGSlice(t *testing.T) {
	// arr := NewArrayFromInterfaceArray(array2)
	arr := NewArray(array1)

	arrResult := []interface{}{}

	for i := -10; i < 10; i++ {
		for j := -10; j < 10; j++ {
			start, end := i, j
			f := arr.Slice(start, end)
			arrResult = append(arrResult, f)
			// fmt.Println(i, j, f)
		}
	}

	// javascript equivalence:
	// var test = () => {
	// array1 = [1, 2, 3, 4, 5, 4, 6]
	// 	ret = []
	// 	for (i = -10; i < 10; i++) {
	// 		for (j = -10; j < 10; j++) {
	// 			start = i;
	// 			end = j
	// 			f = array1.slice(start, end)
	// 			// arrResult = append(arrResult, f)
	// 			ret.push(f)
	// 		}
	// 	}
	// 	return ret
	// }
	// sarr = JSON.stringify(test()).replace(/\,/g, " ")

	fmt.Println(arrResult)
	err := fmt.Sprintf(`%v`, arrResult) != "[[] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [] [2] [2 3] [2 3 4] [2 3 4 5] [2 3 4 5 4] [] [] [2] [2 3] [2 3 4] [2 3 4 5] [2 3 4 5 4] [2 3 4 5 4 6] [2 3 4 5 4 6] [2 3 4 5 4 6] [] [] [] [] [] [] [3] [3 4] [3 4 5] [3 4 5 4] [] [] [] [3] [3 4] [3 4 5] [3 4 5 4] [3 4 5 4 6] [3 4 5 4 6] [3 4 5 4 6] [] [] [] [] [] [] [] [4] [4 5] [4 5 4] [] [] [] [] [4] [4 5] [4 5 4] [4 5 4 6] [4 5 4 6] [4 5 4 6] [] [] [] [] [] [] [] [] [5] [5 4] [] [] [] [] [] [5] [5 4] [5 4 6] [5 4 6] [5 4 6] [] [] [] [] [] [] [] [] [] [4] [] [] [] [] [] [] [4] [4 6] [4 6] [4 6] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [6] [6] [6] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [] [2] [2 3] [2 3 4] [2 3 4 5] [2 3 4 5 4] [] [] [2] [2 3] [2 3 4] [2 3 4 5] [2 3 4 5 4] [2 3 4 5 4 6] [2 3 4 5 4 6] [2 3 4 5 4 6] [] [] [] [] [] [] [3] [3 4] [3 4 5] [3 4 5 4] [] [] [] [3] [3 4] [3 4 5] [3 4 5 4] [3 4 5 4 6] [3 4 5 4 6] [3 4 5 4 6] [] [] [] [] [] [] [] [4] [4 5] [4 5 4] [] [] [] [] [4] [4 5] [4 5 4] [4 5 4 6] [4 5 4 6] [4 5 4 6] [] [] [] [] [] [] [] [] [5] [5 4] [] [] [] [] [] [5] [5 4] [5 4 6] [5 4 6] [5 4 6] [] [] [] [] [] [] [] [] [] [4] [] [] [] [] [] [] [4] [4 6] [4 6] [4 6] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [6] [6] [6] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] []]"

	if err {
		t.Errorf("Test fails, %s", "TestAGSlice")
	}
}

func TestAGSplice(t *testing.T) {
	// arr := NewArrayFromInterfaceArray(array2)

	arrResult := []interface{}{}
	arrRemain := []interface{}{}

	for i := -10; i < 10; i++ {
		for j := -10; j < 10; j++ {
			start, end := i, j
			// array1    = []int{1, 2, 3, 4, 5, 4, 6}
			arr := NewArray(array1)

			f := arr.Splice(start, end)
			arrResult = append(arrResult, f)

			// arr.Splice(start, end)
			arrRemain = append(arrRemain, arr.GetResult())
			// fmt.Println(i, j, f)
		}
	}

	// javascript equivalence:
	// var test = () => {
	// 	ret = []
	// 	for (i = -10; i < 10; i++) {
	// 		for (j = -10; j < 10; j++) {
	// 			array1 = [1, 2, 3, 4, 5, 4, 6]
	// 			start = i;
	// 			end = j
	// 			f = array1.splice(start, end)
	// 			// arrResult = append(arrResult, f)
	// 			ret.push(f)
	// 		}
	// 	}
	// 	return ret
	// }
	// sarr = JSON.stringify(test()).replace(/\,/g, " ")

	fmt.Println("arrResult:", arrResult)
	fmt.Println("arrRemain:", arrRemain)
	err1 := fmt.Sprintf(`%v`, arrResult) != "[[] [] [] [] [] [] [] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [] [] [] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [] [] [] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [] [] [] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [] [] [] [] [] [] [] [2] [2 3] [2 3 4] [2 3 4 5] [2 3 4 5 4] [2 3 4 5 4 6] [2 3 4 5 4 6] [2 3 4 5 4 6] [2 3 4 5 4 6] [] [] [] [] [] [] [] [] [] [] [] [3] [3 4] [3 4 5] [3 4 5 4] [3 4 5 4 6] [3 4 5 4 6] [3 4 5 4 6] [3 4 5 4 6] [3 4 5 4 6] [] [] [] [] [] [] [] [] [] [] [] [4] [4 5] [4 5 4] [4 5 4 6] [4 5 4 6] [4 5 4 6] [4 5 4 6] [4 5 4 6] [4 5 4 6] [] [] [] [] [] [] [] [] [] [] [] [5] [5 4] [5 4 6] [5 4 6] [5 4 6] [5 4 6] [5 4 6] [5 4 6] [5 4 6] [] [] [] [] [] [] [] [] [] [] [] [4] [4 6] [4 6] [4 6] [4 6] [4 6] [4 6] [4 6] [4 6] [] [] [] [] [] [] [] [] [] [] [] [6] [6] [6] [6] [6] [6] [6] [6] [6] [] [] [] [] [] [] [] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [] [] [] [] [] [] [] [2] [2 3] [2 3 4] [2 3 4 5] [2 3 4 5 4] [2 3 4 5 4 6] [2 3 4 5 4 6] [2 3 4 5 4 6] [2 3 4 5 4 6] [] [] [] [] [] [] [] [] [] [] [] [3] [3 4] [3 4 5] [3 4 5 4] [3 4 5 4 6] [3 4 5 4 6] [3 4 5 4 6] [3 4 5 4 6] [3 4 5 4 6] [] [] [] [] [] [] [] [] [] [] [] [4] [4 5] [4 5 4] [4 5 4 6] [4 5 4 6] [4 5 4 6] [4 5 4 6] [4 5 4 6] [4 5 4 6] [] [] [] [] [] [] [] [] [] [] [] [5] [5 4] [5 4 6] [5 4 6] [5 4 6] [5 4 6] [5 4 6] [5 4 6] [5 4 6] [] [] [] [] [] [] [] [] [] [] [] [4] [4 6] [4 6] [4 6] [4 6] [4 6] [4 6] [4 6] [4 6] [] [] [] [] [] [] [] [] [] [] [] [6] [6] [6] [6] [6] [6] [6] [6] [6] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] []]"
	err2 := fmt.Sprintf(`%v`, arrRemain) != "[[1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [2 3 4 5 4 6] [3 4 5 4 6] [4 5 4 6] [5 4 6] [4 6] [6] [] [] [] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [2 3 4 5 4 6] [3 4 5 4 6] [4 5 4 6] [5 4 6] [4 6] [6] [] [] [] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [2 3 4 5 4 6] [3 4 5 4 6] [4 5 4 6] [5 4 6] [4 6] [6] [] [] [] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [2 3 4 5 4 6] [3 4 5 4 6] [4 5 4 6] [5 4 6] [4 6] [6] [] [] [] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 3 4 5 4 6] [1 4 5 4 6] [1 5 4 6] [1 4 6] [1 6] [1] [1] [1] [1] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 4 5 4 6] [1 2 5 4 6] [1 2 4 6] [1 2 6] [1 2] [1 2] [1 2] [1 2] [1 2] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 5 4 6] [1 2 3 4 6] [1 2 3 6] [1 2 3] [1 2 3] [1 2 3] [1 2 3] [1 2 3] [1 2 3] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 4 6] [1 2 3 4 6] [1 2 3 4] [1 2 3 4] [1 2 3 4] [1 2 3 4] [1 2 3 4] [1 2 3 4] [1 2 3 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 6] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [2 3 4 5 4 6] [3 4 5 4 6] [4 5 4 6] [5 4 6] [4 6] [6] [] [] [] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 3 4 5 4 6] [1 4 5 4 6] [1 5 4 6] [1 4 6] [1 6] [1] [1] [1] [1] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 4 5 4 6] [1 2 5 4 6] [1 2 4 6] [1 2 6] [1 2] [1 2] [1 2] [1 2] [1 2] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 5 4 6] [1 2 3 4 6] [1 2 3 6] [1 2 3] [1 2 3] [1 2 3] [1 2 3] [1 2 3] [1 2 3] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 4 6] [1 2 3 4 6] [1 2 3 4] [1 2 3 4] [1 2 3 4] [1 2 3 4] [1 2 3 4] [1 2 3 4] [1 2 3 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 6] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6]]"

	if err1 || err2 {
		t.Errorf("Test fails, %s", "TestAGSlice")
	}
}

func TestAGChunk(t *testing.T) {
	arr := NewArray(array1)
	arrResult, _ := arr.Chunk(2)
	err2 := false
	// if err != nil {
	// 	goto ERR
	// }

	fmt.Println(arrResult)
	err2 = fmt.Sprintf("%v", arrResult) != "[[1 2] [3 4] [5 4] [6]]"

	// ERR:
	if err2 {
		t.Errorf("Test fails, %s", "TestAGChunk")
	}
}

func TestAGConcat(t *testing.T) {
	arr2 := []interface{}{
		44, 55, 66,
		// []interface{}{44, 55, 66},
		[]interface{}{771, 772, []interface{}{8881, 8882}},
	}

	// // js equivalence:
	// array1 = [1, 2, 3, 4, 5, 4, 6]
	// res = array1.concat([44,55,66,[771,772,[8881,8882]]])
	// JSON.stringify(res).replace(/\,/g, " ")

	arr := NewArray(array1)
	arrResult := arr.Concat(arr2)

	fmt.Println(arrResult)
	err := fmt.Sprintf("%v", arrResult) !=
		"[1 2 3 4 5 4 6 44 55 66 [771 772 [8881 8882]]]"

	if err {
		t.Errorf("Test fails, %s", "TestAGConcat")
	}
}

func TestAGCopyWithin(t *testing.T) {
	arr := NewArray(array1)
	arrResult := arr.CopyWithin(3, 2, 4)

	fmt.Println(arrResult)
	err := fmt.Sprintf("%v", arrResult) != "[1 2 3 3 4 4 6]"

	if err {
		t.Errorf("Test fails, %s", "TestAGCopyWithin")
	}
}

/////////////////////// BENCHMARK ////////////////////////

func BenchmarkA(b *testing.B) {
	arr := NewArrayFromInterfaceArray(array2)
	// ln := arr.Length()

	for i := 0; i < b.N; i++ {
		// arr.Join(", ")
		arr.Reduce(func(tot, item interface{}, index int, array []interface{}) interface{} {
			// ii := item.(int)      //* 2
			return tot.(int) + item.(int) // tot.(string) + fmt.Sprintf("%d", ii)
		}, 0)
	}
}

func BenchmarkB(b *testing.B) {

	for i := 0; i < b.N; i++ {
		// arr := NewArrayFromInterfaceArray(array2)
		arr := NewArray(arrayBig)
		// _ = arr.Filter(func(item interface{}, index int, array []interface{}) bool {
		// 	ii := item.(int)
		// 	return ii > 4
		// }).GetResult()

		// arr.Splice(1, 2, "ho", "jaja")
		// arr.Shuffle()
		arr.CopyWithin(3, 12, 34)
	}
}

func TestArray_CopyWithin(t *testing.T) {
	type args struct {
		target int
		start  int
		end    int
	}
	tests := []struct {
		name string
		pa   *Array
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			pa:   NewArray(array1),
			args: args{3, 2, 4},
			want: []interface{}{1, 2, 3, 3, 4, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pa.CopyWithin(tt.args.target, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Array.CopyWithin() = %v, want %v", got, tt.want)
			}
		})
	}
}
