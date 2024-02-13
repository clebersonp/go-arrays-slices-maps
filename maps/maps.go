package main

import "fmt"

// to run this application in terminal: go run maps.go

func main() {
	// map syntax: map[keyType]valueType{}
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Printf("Websites type: %T, content: %v\n", websites, websites)

	// adding a new key-value to maps in Go is so easy, it's dynamic
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Printf("After added a new key-value: %v\n", websites)

	// remove a key-value from map with a built-in delete(map, key) function
	delete(websites, "Google")
	fmt.Printf("After delete a key-value: %v\n", websites)

	// there is no error if you try to delete a key that does not exist in map
	delete(websites, "some-key")
	fmt.Printf("After try delete a key-value that does not exists: %v\n", websites)

}
