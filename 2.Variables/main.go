package main

import "fmt"

func main() {
	// Declare with type and use
	var s string
	s = "Hello"
	fmt.Println(s)

	// Declare and instantiate with type inferred
	i := 10
	fmt.Println(i)

	// Note: variable names are often short in idomatic golang
}
