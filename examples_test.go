package jsarray_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gunawan-pad/go-jsarray"
)

func ExampleNewArray() {
	array1 := []int{1, 2, 3, 4, 5, 4, 6}
	array2 := []interface{}{1, 2, 3, 4, 5, 4, 6}

	// Declaration
	arr := jsarray.NewArray(array1)
	// or
	arr = jsarray.NewArrayFromInterfaceArray(array2)
	// or
	arr = &jsarray.Array{1, 2, 3, 4, 5, 4, 6}

	arr.Sort(func(a, b interface{}) bool {
		return a.(int) > b.(int)
	})

	fmt.Printf("%v", arr.GetResult())
	// Output:
	// [6 5 4 4 3 2 1]
}

// Declaration
func Example_declaration() {
	array1 := []int{1, 2, 3, 4, 5, 4, 6}
	array2 := []interface{}{1, 2, 3, 4, 5, 4, 6}

	var arr *jsarray.Array
	arr = jsarray.NewArray(array1)
	// or
	arr = jsarray.NewArrayFromInterfaceArray(array2)
	// or
	arr = &jsarray.Array{1, 2, 3, 4, 5, 4, 6}
	// or
	arr = (*jsarray.Array)(&array2)
	// or
	arr = (*jsarray.Array)(&[]interface{}{1, 2, 3, 4, 5, 4, 6})

	fmt.Println(arr.GetResult())
	// Output:
	// [1 2 3 4 5 4 6]
}

// This example shows method chaining using function literal
func Example_methodChainingLiteral() {
	arrResult := jsarray.
		NewArray([]int{1, 2, 3, 4, 5, 4, 6}). // initial array is [1 2 3 4 5 4 6]
		Map(func(item interface{}, index int, array []interface{}) interface{} {
			return item.(int) * 2
		}). // [2 4 6 8 10 8 12]
		Filter(func(item interface{}, index int, array []interface{}) bool {
			ii := item.(int)
			return ii > 4
		}).        // [6 8 10 8 12]
		Reverse(). // [12 8 10 8 6]
		Sort(func(a, b interface{}) bool {
			return a.(int) < b.(int)
		}).         // [6 8 8 10 12]
		GetResult() // Get the result array ([]interface{})

	fmt.Println(arrResult) // [6 8 8 10 12]
	// Output:
	// [6 8 8 10 12]
}

// This example shows method chaining using function variables
func Example_methodChainingFuncVar() {
	funcMap := func(item interface{}, index int, array []interface{}) interface{} {
		return item.(int) * 2
	}

	funcFilter := func(item interface{}, index int, array []interface{}) bool {
		return item.(int) > 4
	}

	funcSort := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}

	arrResult := jsarray.NewArray([]int{1, 2, 3, 4, 5, 4, 6}).
		Map(funcMap).
		Filter(funcFilter).
		Reverse().
		Sort(funcSort).
		GetResult()

	fmt.Println(arrResult) // [6 8 8 10 12]
	// Output:
	// [6 8 8 10 12]
}

// This example shows how to use jsarray for processing array in a JSON file
func Example_jSON() {
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
	byt, err := ioutil.ReadFile(`./samples/Rock the 2000's.json`)
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

	// fmt.Println(arr)

	byt, err = json.Marshal(arr)
	if err != nil {
		panic(err)
	}

	outFile := "./samples/testfilter.json"

	// Save result array to json file: "testfilter.json"
	ioutil.WriteFile(outFile, byt, 0777)

}
