# Go JSArray

A utility for bla

# Installing

```
go get -u github.com/gunawan-pad/go-jsarray
```

# Examples

```go
import (
    "fmt"
    "github.com/gunawan-pad/go-jsarray/jsarray"
)

var (
    array1    = []int{1, 2, 3, 4, 5, 4, 6} // int array
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

// Map
arr := jsarray.NewArray(array1)
arrResult := arr.Map(func(item interface{}, index int, array []interface{}) interface{} {
    return item.(int) * 2
}).Get(true).([]interface{})

fmt.Println(arrResult) // [2, 4, 6, 8, 10, 8, 12]

```

Or:

```go
callbackfn := func(item interface{}, index int, array []interface{}) interface{} { return item.(int) * 2 }

arr := jsarray.NewArrayFromInterfaceArray(array2)
arrResult := arr.
    Map(callbackfn).
    Get(true).([]interface{})

fmt.Println(arrResult) // [2, 4, 6, 8, 10, 8, 12]
```

Method chaining:

```go
arrResult := jsarray.NewArray(array1).
    Map(func(item interface{}, index int, array []interface{}) interface{} {
        return item.(int) * 2
    }).
    Filter(func(item interface{}, index int, array []interface{}) bool {
        ii := item.(int)
        return ii > 4
    }).
    Get(false).([]interface{})

fmt.Println(arrResult) // [6 8 10 8 12]
    
```

Sorting string array:

```go
arr := jsarray.NewArray(arrString)
arrResult := arr.Sort(func(a, b interface{}) bool {
    return a.(string) < b.(string)
}).Get(false).([]interface{})

fmt.Println(arrResult) // ["dua", "empat", "empat", "enam", "lima", "satu", "tiga"]

```
