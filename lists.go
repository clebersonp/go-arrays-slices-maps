package main

import "fmt"

func main() {
	// array syntax: declaration and initialization: [num of elements]type{initialize values}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Printf("Price Type: %T, Prices: %v\n", prices, prices)
	fmt.Println(prices[0]) // accessing an array element by position from 0 until length - 1

	// if we create an array without length in declaration, we are creating an empty array,
	// and we can't put any element there later
	var productNames [4]string = [4]string{"A Book"}
	fmt.Println(productNames)
	productNames[2] = "A Carpet" // set an array position value
	fmt.Println(productNames)

	// an empty array can't insert into any element later
	var countryNames []string
	fmt.Println(countryNames)
	//countryNames[0] = "Brazil" // panic: runtime error: index out of range [0] with length 0

}
