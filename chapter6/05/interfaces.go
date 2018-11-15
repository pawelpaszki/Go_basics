package main

import "fmt"

type Animal interface { // custom interface
	Speak() string
}

// below: three custom types implementing the Animal interface
type Dog struct {

}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {

}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cow) Speak() string {
	return "Moo!"
}

type Cow struct {

}


func main() {

	animals := []Animal{Dog{}, Cat{}, Cow{}}

	// for i := range animals {
	// 	fmt.Println(animals[i].Speak())
	// }

	// another for loop example:
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

}