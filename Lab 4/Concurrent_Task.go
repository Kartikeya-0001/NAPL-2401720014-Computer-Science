package main

import "fmt"

func square(ch chan int) {
	fmt.Println("square start")
	ch <- 5 * 5
	fmt.Println("square complete")
}
func cube(ch chan int) {
	fmt.Println("cube start")
	ch <- 3 * 3 * 3
	fmt.Println("cube complete")
}
func fibonacci(ch chan int) {
	fmt.Println("fibonacci start")
	a, b := 0, 1
	for i := 0; i < 6; i++ {
		a, b = b, a+b
	}
	ch <- a
	fmt.Println("fibonacci completed")
}
func main() {
	ch := make(chan int, 3)

	go square(ch)
	go cube(ch)
	go fibonacci(ch)

	var result []int

	for i := 0; i < 3; i++ {
		result = append(result, <-ch)
	}
	fmt.Println("Result:", result)
}
