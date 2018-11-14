package main

import (
	"fmt"
)

// Note: slices are much more powerful than arrays and should be used instead of them
func main() {
	var colors [3]string // items have to be of the same type
	colors[0] = "Red" // no ':' as the type is already inferred
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5]int{1,2,3,4,5}
	fmt.Println(numbers)
	fmt.Println("numbers length is:", len(numbers))
}