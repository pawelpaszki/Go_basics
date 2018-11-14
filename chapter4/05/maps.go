package main

import (
	"fmt"
	"sort"
)

func main() {
	// pretty common to use string for keys
	states := make(map[string]string)
	fmt.Println(states) // map[]

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"

	fmt.Println(states) // map[WA:Washington OR:Oregon CA:California]

	california := states["CA"]
	fmt.Println(california) // California

	// delete an item from a map:
	delete(states, "OR")
	fmt.Println(states) // map[CA:California WA:Washington]

	states["NY"] = "New York"

	// iterate over the map (order not guaranteed):

	for k,v := range states {
		fmt.Printf("%v: %v\n", v, k)
	}

	// map in alphabetical order:

	keys := make([]string, len(states)) // capacity omitted
	i := 0
	for k := range states { // get all of the keys
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	fmt.Println("\nSorted:")
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}