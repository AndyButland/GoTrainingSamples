package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var res int
	res = add(2, 3)
	expected := 5
	if res != expected {
		t.Errorf("Result of addition expected to be %d, got %d.", expected, res)
	}
}

func TestDivide(t *testing.T) {
	var res int
	res, _ = divide(10, 2)
	expected := 5
	if res != expected {
		t.Errorf("Result of division expected to be %d, got %d.", expected, res)
	}

	_, err := divide(10, 0)
	if err == nil {
		t.Errorf("Expected error when dividing by zero.")
	}
}
