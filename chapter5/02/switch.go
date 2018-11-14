package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	rand.Seed(time.Now().Unix())
	// dow := rand.Intn(6) + 1 moved to the start of switch statement
	// fmt.Println("Day:", dow)

	result := ""

	switch dow := rand.Intn(7) + 1; dow { // no explicit break needed
		case 1:
			result = "It's Sunday!"
		case 7:
			result = "It's Saturday!"
		default:
			result = "It's a weekday!"
	}

	fmt.Println(result)

	x := -42
	switch {
		case x < 0:
			result = "less than zero"
			fallthrough // to execute the next cases too..
		case x == 0:
			result = "equal to zero"
		case x > 0:
			result = "greater than zero"
	}

	fmt.Println(result)
}