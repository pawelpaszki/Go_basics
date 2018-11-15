package main

import "fmt"

type Dog struct {
	Breed string
	Weight int
	Sound string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d *Dog) SpeakThreeTimes() { // the "*" means that the value of the object is changed, e.g. pointed is passed and not value
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dog{"Poodle", 37, "Woof"}
	fmt.Println(poodle)
	poodle.Speak()
	
	poodle.Sound = "Arf"
	poodle.Speak()
	
	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()
}
