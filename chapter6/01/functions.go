package main

import (
	"fmt"
)

func main() {
	doSomething()
	sum := addValues(3,5)
	fmt.Println(sum)

	sum = addAllValues(17,23,35,49)
	fmt.Println(sum)
}

// no function overloading in Go ! starting function name with lowercase makes it private to the package
func doSomething() {
	fmt.Println("Doing something!")
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

// add as many params as desired (provided that they are of the same type) 
func addAllValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values) // returns []int (slice not an array), if ints provided as values
	return sum
}