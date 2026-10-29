package main

import "testing"

func TestSqrt(t *testing.T) {
	result, err := calculate("sqrt", 9)
	if err != nil || result != 3 {
		t.Error("Square root test failed")
	}
}

func TestNegative(t *testing.T) {
	_, err := calculate("sqrt", -4)
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestDivideZero(t *testing.T) {
	_, err := calculate("divide", 0)
	if err == nil {
		t.Error("Expected an error")
	}
}