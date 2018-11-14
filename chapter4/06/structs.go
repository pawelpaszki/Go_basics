package main

import (
	"fmt"
)

// struct definition (own custom type)
type Dog struct {
	Breed string// start with uppercase to make those fields public (globally accessible), use single words
	Weight int
}

func main() {
	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle) // {Poodle 34}
	fmt.Printf("%+v\n", poodle) // {Breed:Poodle Weight:34}
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight) // Breed: Poodle
																																		 // Weight: 34
}