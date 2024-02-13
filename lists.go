package main

import "fmt"

// To run this code in terminal: go run lists.go

func main() {
	// To create a dynamic list of slices we just need to omit its length in creation
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1]) // to access the first element of slice
	prices[1] = 9.99         // override the second element of slice(array)
	fmt.Println(prices)

	// to append a new element to slice, Go manage it for us just using the built-in append([]slice, elements...) function
	prices = append(prices, 5.99)
	fmt.Printf("Added element: %v\n", prices)
	// Go has not a built-in function to remove elements from slice(array), we need to use with position feature [:]
	prices = prices[1:] // removing the first (0 - zero) element and returning a new slice(array)
	fmt.Printf("Removed an element: %v\n", prices)
}

//func main() {
//	// array syntax: declaration and initialization: [num of elements]type{initialize values}
//	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
//	fmt.Printf("Price Type: %T, Prices: %v\n", prices, prices)
//	fmt.Println(prices[0]) // accessing an array element by position from 0 until length - 1
//
//	// if we create an array without length in declaration, we are creating an empty array,
//	// and we can't put any element there later
//	var productNames [4]string = [4]string{"A Book"}
//	fmt.Println(productNames)
//	productNames[2] = "A Carpet" // set an array position value
//	fmt.Println(productNames)
//
//	// an empty array can't insert into any element later
//	var countryNames []string
//	fmt.Printf("Country Names: %v\n", countryNames)
//	//countryNames[0] = "Brazil" // panic: runtime error: index out of range [0] with length 0
//
//	// slices: create a slice from an array. Slices is a subset of an array elements
//	// first element is included position and the second one is excluded
//	featuredPrices := prices[1:3]
//	fmt.Printf("Featured Prices Type: %T, elements: %v\n", featuredPrices, featuredPrices)
//
//	// create a slice by variation of its positions
//	// array[1:] => from element of index 1 until the last one
//	// array[:1] => from first element until index 1 (excluded) -1 = 0
//	// array[:] => all elements of an array
//	// array[:len(array)] => all elements of an array, it's variation of array[:]
//	anotherFeaturedPrices := prices[1:]
//	fmt.Println(anotherFeaturedPrices)
//
//	// slice from a slice
//	highLightedPrices := anotherFeaturedPrices[:1]
//	fmt.Println(highLightedPrices)
//
//	fmt.Println("-----------------------")
//	featuredPrices = prices[1:]
//	// slices override the original array value because slices just are pointer of an array range or window of values
//	featuredPrices[0] = 199.99
//	highLightedPrices = featuredPrices[:1]
//	fmt.Println(highLightedPrices)
//	fmt.Println(prices)
//	fmt.Println(len(highLightedPrices), cap(highLightedPrices))
//
//	// highLightedPrices has only a pointer of one element (len()) but Go memorize the original capacity (cap())
//	// therefore we can create another slices looking forwards in capacity and create a slices "greater than" the last one
//	highLightedPrices = highLightedPrices[:3]
//	fmt.Println(highLightedPrices)
//	fmt.Println(len(highLightedPrices), cap(highLightedPrices))
//}
