package library

import (
	"testing"
)

func TestAdditon(t *testing.T) {
	if Additon(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3")
	}
	if Additon(-1, -2) != -3 {
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}

func TestSubtraction(t *testing.T) {
	if Subtraction(2, 1) != 1 {
		t.Error("Expected 2 (-) 1 to equal 1")
	}
	if Subtraction(-2, -1) != -1 {
		t.Error("Expected -2 (-) -1 to equal -1")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(2, 2) != 4 {
		t.Error("Expected 2 (*) 2 to equal 4")
	}
	if Multiplication(2, -2) != -4 {
		t.Error("Expected 2 (*) -2 to equal -4")
	}
}

func TestDivision(t *testing.T) {
	if Division(2, 2) != 1 {
		t.Error("Expected 2 (/) 2 to equal 1")
	}
	if Division(-4, 2) != -2 {
		t.Error("Expected -4 (/) 2 to equal -2")
	}
}
