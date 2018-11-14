package main

import (
	"strconv"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	// var s string
	// fmt.Scanln(&s) // will only get "one", when supplied with "one two three"
	// fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)


	fmt.Print("Enter a number: ")
	str, _ = reader.ReadString('\n') // removed ":" -> cannot redeclare a variable
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64) // 64 -> type of number
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", f)
	}
}