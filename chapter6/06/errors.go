package main

import (
	"fmt"
	"os"
	"errors"
)

func main() {
	f, err := os.Open("textfile.txt")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}

	///
	myerror := errors.New("My error string") // create new error
	fmt.Println(myerror)
	///

	attendance := map[string]bool {
		"Ann": true,
		"Rob": true}

	attended, ok := attendance["Mike"] // if there is Mike in the map, the ok will be true, Mike will have value of true in this case
	if ok { // if not ok -> "Mike" value not present in the map
		fmt.Println("Mike attended?", attended)
	} else {
		fmt.Println("No info for Mike")
	}
}