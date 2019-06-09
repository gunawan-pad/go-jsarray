package jsarray

import (
	"fmt"
	"testing"
)

var (
	array1    = []int{1, 2, 3, 4, 5, 4, 6}         // int array
	array2    = []interface{}{1, 2, 3, 4, 5, 4, 6} // interface array
	arrString = []string{
		"satu", "dua", "tiga", "empat", "lima", "empat", "enam",
	}
	arrayNested = []interface{}{
		12, 23,
		[]interface{}{31, 32, 33, []interface{}{331, 332}, 34, 35},
		41, 41,
		[]interface{}{51, 52, 53},
	}
)

func CompareArray(arrResult, arrCompare []interface{}) bool {
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
	var arrCompare = []interface{}{3, 3, 3, 3, 3, 3, 3, 5, 5, -1, 3, 3, 3, 3, 5, 5, -1, -1, -1, -1}
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
	// TODO: TestAGSlice

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
	// sarr = test().replace(/\,/g, " ")

	fmt.Println(arrResult)
	err := fmt.Sprintf(`%v`, arrResult) != "[[] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [] [2] [2 3] [2 3 4] [2 3 4 5] [2 3 4 5 4] [] [] [2] [2 3] [2 3 4] [2 3 4 5] [2 3 4 5 4] [2 3 4 5 4 6] [2 3 4 5 4 6] [2 3 4 5 4 6] [] [] [] [] [] [] [3] [3 4] [3 4 5] [3 4 5 4] [] [] [] [3] [3 4] [3 4 5] [3 4 5 4] [3 4 5 4 6] [3 4 5 4 6] [3 4 5 4 6] [] [] [] [] [] [] [] [4] [4 5] [4 5 4] [] [] [] [] [4] [4 5] [4 5 4] [4 5 4 6] [4 5 4 6] [4 5 4 6] [] [] [] [] [] [] [] [] [5] [5 4] [] [] [] [] [] [5] [5 4] [5 4 6] [5 4 6] [5 4 6] [] [] [] [] [] [] [] [] [] [4] [] [] [] [] [] [] [4] [4 6] [4 6] [4 6] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [6] [6] [6] [] [] [] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [] [1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5] [1 2 3 4 5 4] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [1 2 3 4 5 4 6] [] [] [] [] [] [2] [2 3] [2 3 4] [2 3 4 5] [2 3 4 5 4] [] [] [2] [2 3] [2 3 4] [2 3 4 5] [2 3 4 5 4] [2 3 4 5 4 6] [2 3 4 5 4 6] [2 3 4 5 4 6] [] [] [] [] [] [] [3] [3 4] [3 4 5] [3 4 5 4] [] [] [] [3] [3 4] [3 4 5] [3 4 5 4] [3 4 5 4 6] [3 4 5 4 6] [3 4 5 4 6] [] [] [] [] [] [] [] [4] [4 5] [4 5 4] [] [] [] [] [4] [4 5] [4 5 4] [4 5 4 6] [4 5 4 6] [4 5 4 6] [] [] [] [] [] [] [] [] [5] [5 4] [] [] [] [] [] [5] [5 4] [5 4 6] [5 4 6] [5 4 6] [] [] [] [] [] [] [] [] [] [4] [] [] [] [] [] [] [4] [4 6] [4 6] [4 6] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [6] [6] [6] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] []]"

	if err {
		t.Errorf("Test fails, %s", "TestAGSlice")
	}
}

/////////////////////// BENCHMARK ////////////////////////

func BenchmarkJoin(b *testing.B) {
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

func BenchmarkFilter(b *testing.B) {

	for i := 0; i < b.N; i++ {
		arr := NewArrayFromInterfaceArray(array2)
		_ = arr.Filter(func(item interface{}, index int, array []interface{}) bool {
			ii := item.(int)
			return ii > 4
		}).GetResult()
	}
}
