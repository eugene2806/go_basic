package main

import "fmt"

//1. Данный метод явно связан только с University

func (u *University) FullUniversityInfo() {
	fmt.Printf("Uni name %s and City %s\n", u.name, u.city)
}

//2. Но в структуру Professor встроены поля University

type Professor struct {
	fullName string
	age      int
	University
}

type University struct {
	city string
	name string
}

func main() {
	p := Professor{
		fullName: "Boris Bobroff",
		age:      150,
		University: University{
			city: "Moscow",
			name: "BGPTU",
		},
	}
	p.FullUniversityInfo()
}
