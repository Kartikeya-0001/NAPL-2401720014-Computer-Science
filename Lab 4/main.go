package main

import (
	"errors"
	"fmt"
	"math"
)

func calculate(choice string, num float64) (float64, error) {
	if choice == "" {
		return 0, errors.New("input is empty")
	}

	if choice == "divide" && num == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	if choice == "sqrt" && num < 0 {
		return 0, errors.New("negative number")
	}

	if choice == "sqrt" {
		return math.Sqrt(num), nil
	}

	if choice == "divide" {
		return 10 / num, nil
	}

	return 0, errors.New("wrong choice")
}

func main() {
	result, err := calculate("sqrt", -4)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}