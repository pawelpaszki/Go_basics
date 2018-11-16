package main

import (
	"io/ioutil"
	"fmt"
	"io"
	"os"
)

func main() {

	content := "Hello from Go!"

	file, err := os.Create("./fromString.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("All done with file of %v characters", ln)

	// byte 
	bytes := []byte(content)
	ioutil.WriteFile("./fromBytes.txt", bytes, 0644)

}

func checkError(err error) {
	if err != nil {
		panic(err) // causes the application to stop (panic)
	}
}