package main

import (
	"io/ioutil"
	"fmt"
)

func main() {

	filename := "./hello.txt"

	content, err := ioutil.ReadFile(filename)
	checkError(err)

	result := string(content) // convert byte array to string

	fmt.Println("Read from file:", result)
}

func checkError(err error) {
	if err != nil {
		panic(err) // causes the application to stop (panic)
	}
}