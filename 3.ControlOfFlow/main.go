package main

import "fmt"

func main() {
	// 3 ways to count to five
	// - as you'd expect
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("-")

	// - use 'for' for 'while'
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println("-")

	// - with no clause for 'while(true)'
	i = 1
	for {
		fmt.Println(i)
		i++
		if i > 5 {
			break
		}
	}
	fmt.Println("-")

	// If/else with fizzbuzz
	// - note: no ternary operator ( ? : )
	for i := 1; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
