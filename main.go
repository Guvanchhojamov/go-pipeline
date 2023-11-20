package main

import "fmt"

var version = "dev"
var i int

func main() {
	fmt.Printf("Version: %s\n", version)
	fmt.Println(hello())
}

func hello() string {
	return "Hello Golang!"
}
