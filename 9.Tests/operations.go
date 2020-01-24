package calculator

import "errors"

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("attempt to divide by zero")
	}

	return x / y, nil
}
