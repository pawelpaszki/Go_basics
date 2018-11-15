package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Close the file") // wait until everything else is finished and then call this function
	fmt.Println("Open the file")

	defer fmt.Println("Statement 1")

	defer fmt.Println("Statement 2")

	myFunc()

	defer fmt.Println("Statement 3")

	defer fmt.Println("Statement 4")

	defer fmt.Println("Undeferred statement")

	x := 1000
	defer fmt.Println("Value of x:", x)
	x++
	fmt.Println("Value of x:", x)
}

func myFunc() {
	defer fmt.Println("Deferred in the function")
	fmt.Println("Not deferred in the function")
}