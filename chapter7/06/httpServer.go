package main

import (
	"fmt"
	"net/http"
)

type Hello struct{} 

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) { // r -> reference // the types have to be the same
	fmt.Fprint(w, "<h1>Hello from the Go web server!</h1>") // prints to a writer object
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	checkError(err)
}

func checkError(err error) {

}