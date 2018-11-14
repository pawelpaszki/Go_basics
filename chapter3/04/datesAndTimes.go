package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	now := time.Now()
	fmt.Printf("Now is: %s\n", now)
	
	fmt.Println("The month is:", t.Month())
	fmt.Println("The month is:", t.Day())
	fmt.Println("The month is:", t.Weekday())

	tomorrow := t.AddDate(0,0,1)
	fmt.Printf("Tomorrow is: %v, %v, %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	longFormat := "Monday, January 2, 2006"
	fmt.Println("tomorrow is:", tomorrow.Format(longFormat))

	shortFormat := "1/2/06"
	fmt.Println("tomorrow is:", tomorrow.Format(shortFormat))
}