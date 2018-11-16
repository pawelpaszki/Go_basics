package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close() // close the body after completing reading the request

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes) // convert byte array to string 
	fmt.Print(content)

}