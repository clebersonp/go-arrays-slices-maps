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

}
