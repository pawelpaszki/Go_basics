package main

import (
	"fmt"
)

func main() {

	n1, l1 := FullName("Pawel", "Paszki")
	fmt.Printf("Full name: %v, number of chars: %v\n", n1, l1) // %v -> verb

	n1, l1 = FullNameNakedReturn("John", "Doe") // note -> no ':' here
	fmt.Printf("Full name: %v, number of chars: %v\n", n1, l1) // %v -> verb
}

func FullName(f, l string) (string, int) { // note the return types wrappped inside parentheses
	full := f + " " + l
	length := len(full)
	return full, length
}

func FullNameNakedReturn(f, l string) (full string, length int) { // return values called here -> naked return below
	full = f + " " + l // note -> no ':' here
	length = len(full) // note -> no ':' here
	return // not explicitly stating which values to return
}