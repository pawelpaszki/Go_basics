package main

import (
	"fmt"
	"sort"
)

// slice in go -> abstraction layer sitting on top of array
func main() {
	var colors = []string{"Red", "Green", "Blue"} // no number inside the "[]" -> slice and not an array
	fmt.Println(colors)
	
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)]) // len(colors) optional <- this removes the first item
	fmt.Println(colors)

	colors = append(colors[1:]) // remove first item
	fmt.Println(colors)

	colors = append(colors[:len(colors) - 1]) // remove last item
	fmt.Println(colors)

	numbers := make([]int, 5, 5) // create a slice with init size of 5 and capacity also 5
	numbers[0] = 5
	numbers[1] = 4
	numbers[2] = 6
	numbers[3] = 7
	numbers[4] = 2
	//numbers[5] = 6 // index out of range
	fmt.Println(numbers)

	numbers = append(numbers, 0)
	fmt.Println(numbers)
	fmt.Println(cap(numbers)) // capacity -> 10!

	sort.Ints(numbers)
	fmt.Println(numbers)
}