package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p *Person) ReadInput() {
	fmt.Print("Enter Name: ")
	fmt.Scan(&p.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&p.Age)

	fmt.Print("Enter Job: ")
	fmt.Scan(&p.Job)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&p.Salary)
}

func (p Person) Display() {
	fmt.Println("\nPerson Details")
	fmt.Println("Name :", p.Name)
	fmt.Println("Age :", p.Age)
	fmt.Println("Job :", p.Job)
	fmt.Println("Salary :", p.Salary)
}

func main() {
	var p1, p2 Person

	fmt.Println("Enter Person 1 Details:")
	p1.ReadInput()
	p1.Display()

	fmt.Println("\nEnter Person 2 Details:")
	p2.ReadInput()
	p2.Display()
}
