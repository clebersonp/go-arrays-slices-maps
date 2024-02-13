package main

import "fmt"

func main() {
	// array syntax: declaration and initialization: [num of elements]type{initialize values}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Printf("Price Type: %T, Prices: %v\n", prices, prices)
}
