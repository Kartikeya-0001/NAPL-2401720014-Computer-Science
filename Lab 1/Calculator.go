package main

import "fmt"

func main() {
	var choice int

	for {
		fmt.Println("\n--- Simple Calculator ---")
		fmt.Println("1. Integer Calculation")
		fmt.Println("2. Float Calculation")
		fmt.Println("3. Exit")

		fmt.Print("Enter your choice (1-3): ")
		_, err := fmt.Scan(&choice)

		// Validate input
		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("Invalid input! Please enter a number between 1 and 3.")

			var clear string
			fmt.Scanln(&clear)
			continue
		}

		if choice == 3 {
			fmt.Println("Calculator closed.")
			break
		}

		if choice == 1 {
			var a, b int

			fmt.Print("Enter first integer: ")
			fmt.Scan(&a)

			fmt.Print("Enter second integer: ")
			fmt.Scan(&b)

			fmt.Println("\nAddition:", a+b)
			fmt.Println("Subtraction:", a-b)
			fmt.Println("Multiplication:", a*b)

		} else if choice == 2 {
			var x, y float64

			fmt.Print("Enter first decimal number: ")
			fmt.Scan(&x)

			fmt.Print("Enter second decimal number: ")
			fmt.Scan(&y)

			fmt.Println("\nAddition:", x+y)
			fmt.Println("Subtraction:", x-y)
			fmt.Println("Multiplication:", x*y)
		}
	}
}
