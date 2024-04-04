package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

// 1. Методы - это функции привязанные к определенным структурам
// func (s Struct) MethodName(parameters type) type {}
//	     Receiver - получатель метода

func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Printf("Salary: %d %s\n", e.salary, e.currency)
}

func main() {
	emp := Employee{
		"Bob",
		"Senior Developer",
		3000,
		"USD",
	}

	emp.DisplayInfo()
}
