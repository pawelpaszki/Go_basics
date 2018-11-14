package main

import (
	"fmt"
)

func main() {
	// var x float64  = 0
	var result string

	if x := 42; x < 0 { // x only available inside the if statement
		result = "less than zero"
	} else if x == 0 {
		result = "equal to zero"
	} else {
		result = "greater than zero"
	}
	
	fmt.Println("Result:", result)
}