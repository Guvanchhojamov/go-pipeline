package main

import "fmt"

var version = "dev"
var i int

func main() {

	// This is main function

	fmt.Printf("Version: %s\n", version)
	fmt.Println(hello())
}

func hello() string {
	return "Hello Golang!"
}
