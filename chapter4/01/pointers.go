package main

import (
	"fmt"
)

func main() {
	var p *int // pointer
	
	if p != nil {
		fmt.Println("Value of p:", *p) // returns an error (nil)
	} else {
		fmt.Println("p is nil")
	}
	
	var v int = 42
	p = &v // connect pointer to this variable

	if p != nil {
		fmt.Println("Value of p:", *p) // returns 42
	} else {
		fmt.Println("p is nil")
	}

	var value1 float64 = 42.13
	pointer1 := &value1

	fmt.Println("value1:", *pointer1)

	*pointer1  = *pointer1 / 31
	fmt.Println("pointer1:", *pointer1)
	fmt.Println("value1:", value1)
}