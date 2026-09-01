package main

import (
	"fmt"

	"Main.go/myutil"
)

func main() {

	text := "New Age Programming Lab"

	fmt.Println("Original String:", text)
	fmt.Println("Reverse:", myutil.Reverse(text))
	fmt.Println("Vowel Count:", myutil.CountVowels(text))

	fmt.Println("Factorial of 5:", myutil.Factorial(5))
	fmt.Println("2 raised to power 5:", myutil.Power(2, 5))
}
