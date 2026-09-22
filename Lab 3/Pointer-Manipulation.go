package main

import "fmt"

type Student struct {
	Name  string
	Marks float64
	Age   int
}

func modifyAge(age *int) {
	*age = 20
}

func main() {

	age := 18
	ptr := &age

	fmt.Println("Value of age:", age)
	fmt.Println("Address of age:", ptr)
	fmt.Println("Value using pointer:", *ptr)

	fmt.Println("\nBefore modification:", age)

	modifyAge(&age)

	fmt.Println("After modification:", age)

	s := new(Student)

	s.Name = "Kartikeya"
	s.Marks = 88.8
	s.Age = 20

	fmt.Println("\nStudent details:")
	fmt.Println("Name:", s.Name)
	fmt.Println("Marks:", s.Marks)
	fmt.Println("Age:", s.Age)

	s.Marks = 99.9

	fmt.Println("\nAfter modifying marks:")
	fmt.Println("Marks:", s.Marks)
}
