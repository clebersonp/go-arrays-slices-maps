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
	fmt.Printf("Country Names: %v\n", countryNames)
	//countryNames[0] = "Brazil" // panic: runtime error: index out of range [0] with length 0

	// slices: create a slice from an array. Slices is a subset of an array elements
	// first element is included position and the second one is excluded
	featuredPrices := prices[1:3]
	fmt.Printf("Featured Prices Type: %T, elements: %v\n", featuredPrices, featuredPrices)

	// create a slice by variation of its positions
	// array[1:] => from element of index 1 until the last one
	// array[:1] => from first element until index 1 (excluded) -1 = 0
	// array[:] => all elements of an array
	// array[:len(array)] => all elements of an array, it's variation of array[:]
	anotherFeaturedPrices := prices[1:]
	fmt.Println(anotherFeaturedPrices)

	// slice from a slice
	highLightedPrices := anotherFeaturedPrices[:1]
	fmt.Println(highLightedPrices)

}
