package main

import (
	"fmt"
	"stringutil"
)

func main() {

	n1, l1 := stringutil.FullName("Pawel", "Paszki")
	fmt.Printf("Full name: %v, number of chars: %v\n", n1, l1) // %v -> verb

	n1, l1 = stringutil.FullNameNakedReturn("John", "Doe") // note -> no ':' here
	fmt.Printf("Full name: %v, number of chars: %v\n", n1, l1) // %v -> verb
}