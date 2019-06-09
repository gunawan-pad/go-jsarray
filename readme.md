# Go-JSArray

`Go-JSArray` is a library for processing array like in javascript.

  - Support any data type (using `interface{}`)
  - Support method chaining

Supported methods:
- `Every`
- `Fill`
- `Filter`
- `Find`
- `FindIndex`
- `Flat`
- `ForEach`
- `Get`
- `Includes`
- `IndexOf`
- `Join`
- `LastIndexOf`
- `Length`
- `Map`
- `Pop`
- `Push`
- `Reduce`
- `ReduceRight`
- `Reverse`
- `Shift`
- `Slice`
- `Some`
- `Sort`
- `Unshift`

## Installation
```
go get -u github.com/gunawan-pad/go-jsarray
```

## Todos

 - Write MORE Tests
 - Generate code for another data type (string, int etc.) for better performancs
 
## Examples

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

### Method chaining

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

### Sorting string array:

```go
arr := jsarray.NewArray(arrString)
arrResult := arr.Sort(func(a, b interface{}) bool {
    return a.(string) < b.(string)
}).Get(false).([]interface{})

fmt.Println(arrResult) // ["dua", "empat", "empat", "enam", "lima", "satu", "tiga"]

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
		Name           string     //`json:"name"`
		Image          string     `json:"image"`
		Href           string     `json:"href"`
		FollowersTotal int64      `json:"followersTotal"`
		Type           string     //`json:"type"`
		ID             string     //`json:"id"`
		Tracks         []SongInfo //`json:"tracks"`
	}

	file := `Rock the 2000's.json`
	var data Playlist

	byt, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byt, &data)
	if err != nil {
		panic(err)
	}

	arr := jsarray.NewArray(data.Tracks).
		// filter tracks by artist's name started with S
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
		Get(false).([]interface{})

	fmt.Println(arr)
	byt, _ = json.Marshal(arr)
	ioutil.WriteFile("testfilter.json", byt, 0777)
}
```

## License
This work is published under the MIT license.
Please see the `LICENSE` file for details.
