package main

import "fmt"

func main() {
	// ARRAYS are fixed size and must contain the same type

	// Declaring an empty array fills with the "zero-value" of the type.
	var a [3]int
	fmt.Println(a)

	// Set values
	a[0] = 99
	a[1] = 66
	a[2] = 33
	fmt.Println(a)

	// Same array with short-hand
	a2 := [3]int{99, 66, 33}
	fmt.Println(a2)

	// Same array with short-hand and compiler figuring out length
	a3 := [...]int{99, 66, 33}
	fmt.Println(a3)

	// Arrays are VALUE types, so changes to copies (or when passed as arguments to functions)
	// don't change original
	a4 := a3
	a4[0] = 100
	fmt.Println(a3, a4)

	// Length of array
	fmt.Println(len(a))

	// Iterate using for
	for i := 0; i < len(a); i++ {
		fmt.Printf("Element %d of array is %d\n", i, a[i])
	}

	// Or range
	for i, v := range a {
		fmt.Printf("Element %d of array is %d\n", i, v)
	}

	// - we can ignore the index if we don't need it:
	for _, v := range a {
		fmt.Printf("Element value is %d\n", v)
	}

	// SLICES are wrappers around arrays that are unboarded
	// Explicit underlying array
	s := a[:]
	fmt.Println(s)

	// Implicit underlying array
	s2 := []int{1, 2, 3}
	fmt.Println(s2)

	// Change slice changes underlying array
	for i := range s {
		s[i]++
	}
	fmt.Println(s, a)

	// Can append to a slice
	s = append(s, 123)
	fmt.Println(s)
}
