package calculator

import "errors"

func Add(x int, y int) int {
	return x + y
}

func Subtract(x int, y int) int {
	return x - y
}

func Multiply(x int, y int) int {
	return x * y
}

func Divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("attempt to divide by zero")
	}

	return performSafeDivide(x, y), nil
}

func performSafeDivide(x int, y int) int {
	return x / y
}
