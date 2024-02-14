package main

import "fmt"

// Go feature called "type alias"
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {

	// function make to create a slice of string with initial length of 2 and a capacity of max 5 elements
	// make is used when you now about the length of slice, and you can work with position
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Allan")
	userNames = append(userNames, "Mark")

	fmt.Println(userNames)

	fmt.Println("--------------------------")
	// unlike make with slices, with map has only one int argument for initial capacity
	courseRatings := make(floatMap, 4)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()
}
