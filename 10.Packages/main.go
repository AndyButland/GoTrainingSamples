package main

import (
	"fmt"

	"./calculator"
)

func main() {
	var x, y, result int
	x = 8
	y = 2

	// Can call exposed methods (those with PascalCase)
	result = calculator.Add(x, y)
	fmt.Printf("%d plus %d is %d\n", x, y, result)

	result, _ = calculator.Divide(x, y)
	fmt.Printf("%d divided by %d is %d\n", x, y, result)

	// But not private ones (those with camelCase)
	//result, _ = calculator.performSafeDivide(x, y)

	// Error handling
	var err error
	y = 0
	result, err = calculator.Divide(x, y)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("%d divided by %d is %d\n", x, y, result)
	}
}
