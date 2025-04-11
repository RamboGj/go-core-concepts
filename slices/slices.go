package slices

import "fmt"

/*
Slice vs Array
Array is a length defined variable.
Slice is an structure that has a pointer for an array that can have dynamic values and length.
*/

var moviesInDB = []string{
	"Filme 1",
	"Filme 2",
	"Filme 3",
	"Filme 4",
	"Filme 5",
	"Filme 6",
	"Filme 7",
	"Filme 8",
	"Filme 9",
	"Filme 10",
	"Filme 11",
}

func main() {
	IdsFromAPI := []int{1,2,3,4,5,6,7,8,9,10}
	// make() creates a slice without creating new memory allocations as the slice grows. type, len() and cap()
	movies := make([]string, 0, 10)

	someIds := IdsFromAPI[0:2:2] // A slice with the cap defined 
	fmt.Printf("some ids -> %v", someIds)
	
	for _, id := range IdsFromAPI {
		movie := moviesInDB[id]
		fmt.Println(cap(movies))
		movies = append(movies, movie)
	}
	
	fmt.Println(movies)
	
	// sliceRoot := []int{1,2,3}
	// fmt.Println(sliceRoot, len(sliceRoot), cap(sliceRoot))
	
	// arr := [5]int{1,2,3,4,5}
	// slice := arr[1:4]
	// arr[2] = 15
	// slice[0] = 123
	// fmt.Println(arr)
	// fmt.Println(slice)
}
