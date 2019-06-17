![golang](https://img.shields.io/badge/Go-language-blue.svg?logo=go)
[![MIT](https://img.shields.io/badge/license-MIT-orange.svg)](https://github.com/gunawan-pad/go-jsarray/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/gunawan-pad/go-jsarray)](https://goreportcard.com/report/github.com/gunawan-pad/go-jsarray)

# Go-JSArray

`Go-JSArray` is a library for processing golang array like in javascript.

  - Support array of any data type (using internal array of `interface{}`)
  - Support method chaining

Supported methods:
- `Chunk` -> [PHP Reference](https://www.php.net/manual/en/function.array-chunk.php)
- `Concat` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/concat)
- `CopyWithin` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/copywithin)
- `Every` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/every)
- `Fill` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/fill)
- `Filter` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/filter)
- `Find` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/find)
- `FindIndex` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/findindex)
- `Flat` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/flat)
- `ForEach` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/foreach)
- `Includes` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/includes)
- `IndexOf` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/indexof)
- `Join` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/join)
- `LastIndexOf` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/lastindexof)
- `Length` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/length)
- `Map` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/map)
- `Pop` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/pop)
- `Push` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/push)
- `Reduce` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/reduce)
- `ReduceRight` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/reduceright)
- `Reverse` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/reverse)
- `Shift` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/shift)
- `Shuffle` -> [PHP Reference](https://www.php.net/manual/en/function.shuffle.php)
- `Slice` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/slice)
- `Some` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/some)
- `Sort` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/sort)
- `Splice` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/splice)
- `Split` -> [PHP Reference](https://www.php.net/manual/en/function.array-chunk.php)
- `Unique` -> [PHP Reference](https://www.php.net/manual/en/function.array-unique.php)
- `Unshift` -> [JS Reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/unshift)



## Installation
```
go get github.com/gunawan-pad/go-jsarray
```

## Examples

```go
package main

import (
    "fmt"
    "github.com/gunawan-pad/go-jsarray/jsarray"
)

var (
	array1    = []int{1, 2, 3, 4, 5, 4, 6} // int array
	array2    = []interface{}{1, 2, 3, 4, 5, 4, 6} // interface array
	
	arrString = []string{
		"satu", "dua", "tiga", "empat", "lima", "empat", "enam",
	} // string array

	arrayNested = []interface{}{
		12, 23,
		[]interface{}{31, 32, 33, []interface{}{331, 332}, 34, 35},
		41, 41,
		[]interface{}{51, 52, 53},
	}
)

func main() {
	// Map, using function literal
	arr := jsarray.NewArray(array1)
	arr = arr.Map(func(item interface{}, index int, array []interface{}) interface{} {
		return item.(int) * 2
	})
	arrResult := arr.GetResult()

	fmt.Println(arrResult) // [2 4 6 8 10 8 12]
}

```

Or using function variable:

```go
callbackfn := func(item interface{}, index int, array []interface{}) interface{} { return item.(int) * 2 }

arr := jsarray.NewArrayFromInterfaceArray(array2)
arrResult := arr.Map(callbackfn).GetResult()

fmt.Println(arrResult) // [2 4 6 8 10 8 12]
```

### Method chaining

```go
arrResult := jsarray.NewArray(array1). // initial array is [1 2 3 4 5 4 6]
	Map(func(item interface{}, index int, array []interface{}) interface{} {
		return item.(int) * 2 
	}). // [2 4 6 8 10 8 12]
	Filter(func(item interface{}, index int, array []interface{}) bool {
		ii := item.(int)
		return ii > 4 
	}). // [6 8 10 8 12]
	Reverse(). // [12 8 10 8 6]
	Sort(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}). // [6 8 8 10 12]
	GetResult() // Get the result array ([]interface{})

fmt.Println(arrResult) // [6 8 8 10 12]
    
```
### Method chaining, cleaner version

```go
funcMap := func(item interface{}, index int, array []interface{}) interface{} {
	return item.(int) * 2
}

funcFilter := func(item interface{}, index int, array []interface{}) bool {
	ii := item.(int)
	return ii > 4
}

funcSort := func(a, b interface{}) bool {
	return a.(int) < b.(int)
}

arrResult := jsarray.NewArray(array1).
	Map(funcMap).
	Filter(funcFilter).
	Reverse().
	Sort(funcSort).
	GetResult()

fmt.Println(arrResult) // [6 8 8 10 12]
```

### Sorting string array:

```go
arr := jsarray.NewArray(arrString)
arrResult := arr.Sort(func(a, b interface{}) bool {
    return a.(string) < b.(string)
}).GetResult()

fmt.Println(arrResult) // [dua empat empat enam lima satu tiga]

```

### Calculating sum of an array 
```go
// array1    = []int{1, 2, 3, 4, 5, 4, 6} 

array1Sum := jsarray.NewArray(array1).
		Reduce(func(tot, item interface{}, index int, array []interface{}) interface{} {
			return tot.(int) + item.(int)
		}, 0) // 25

fmt.Printf("Sum of array1 is %d\n", array1Sum) // Sum of array1 is 25
```

### Processing JSON file

```go
func TestJSArrayJSONFile() {
	type SongInfo struct {
		Album  string `json:"album"`
		Song   string `json:"song"`
		Href   string `json:"href"`
		ID     string `json:"id"`
		Artist string `json:"artist"`
	}
	type Playlist struct {
		Name           string     `json:"name"`
		Image          string     `json:"image"`
		Href           string     `json:"href"`
		FollowersTotal int64      `json:"followersTotal"`
		Type           string     `json:"type"`
		ID             string     `json:"id"`
		Tracks         []SongInfo `json:"tracks"`
	}

	var data Playlist

	// Open and read json file
	byt, err := ioutil.ReadFile(`Rock the 2000's.json`)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byt, &data)
	if err != nil {
		panic(err)
	}

	arr := jsarray.NewArray(data.Tracks).
		// filter tracks by artist's name started with character 'S'
		Filter(func(item interface{}, index int, array []interface{}) bool {
			artist := item.(SongInfo).Artist
			return strings.HasPrefix(artist, "S")
		}).
		// sort tracks by song title ascending
		Sort(func(a, b interface{}) bool {
			sia := a.(SongInfo)
			sib := b.(SongInfo)
			return sia.Song < sib.Song
		}).
		GetResult() // Get result array

	fmt.Println(arr)
	byt, err = json.Marshal(arr)
	if err != nil {
		panic(err)
	}

	// Save result array to json file: "testfilter.json"
	ioutil.WriteFile("testfilter.json", byt, 0777)
}

// Execute the function
TestJSArrayJSONFile()
```

## License
Copyright © 2019 Gunawan Pad. 
This work is published under the MIT license. 
Please see the [`LICENSE`](LICENSE) file for details.
