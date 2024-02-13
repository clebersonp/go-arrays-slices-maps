package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

type Product struct {
	id    string
	title string
	price float64
}

func New(title string, price float64) Product {
	return Product{
		id:    uuid.NewV4().String(),
		title: title,
		price: price,
	}
}

// To run this code in terminal: go run practice.go
func main() {

	// 1)
	hobbies := [3]string{"Read", "Study", "Travel"}
	fmt.Printf("Hobbies: %v\n", hobbies)

	// 2)
	fmt.Printf("First hobbies element: %v\n", hobbies[0])
	fmt.Printf("Second and third element: %v\n", hobbies[1:3])

	// 3)
	hobbiesSlice := hobbies[:2]
	fmt.Printf("Hobbies Slice: %v, length: %v, capacity: %v\n", hobbiesSlice, len(hobbiesSlice), cap(hobbiesSlice))
	hobbiesSlice = hobbies[0:2]
	fmt.Printf("Hobbies Slice: %v, length: %v, capacity: %v\n", hobbiesSlice, len(hobbiesSlice), cap(hobbiesSlice))

	// 4)
	hobbiesSlice = hobbiesSlice[1:3]
	fmt.Printf("HobbiesSlice: %v\n", hobbiesSlice)

	// 5)
	courseGoals := []string{"Promotion", "Experience"}
	fmt.Printf("Course Goals: %v\n", courseGoals)

	// 6)
	courseGoals[1] = "Skill"
	courseGoals = append(courseGoals, "Knowledge")
	fmt.Printf("Course Goals: %v\n", courseGoals)

	// 7)
	products := []Product{
		New("Notebook", 1000.98),
		New("Monitor", 350.78),
	}
	fmt.Println(products)
	products = append(products, New("Mouse", 15.24))
	fmt.Println(products)
}
