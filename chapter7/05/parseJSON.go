package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"math/big"
)

type Tour struct {
	Name, Price string
}

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

	tours := toursFromJSON(content)

	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {
	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close() // close the body after completing reading the request

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes) // convert byte array to string 
}

func toursFromJSON(content string) []Tour {
	tours := make([]Tour, 0, 20)
	
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()// remove the initial bracket from the json response
	checkError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}

	return tours
}