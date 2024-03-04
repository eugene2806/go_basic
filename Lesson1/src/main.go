package main

import (
	"fmt"
	"os"
)

func main() {
	//Декларирование переменных
	var age int
	fmt.Println("My age is", age)

	//Декларирование и инициализация пользовательским значением
	var height int = 183
	fmt.Println("My height is", height)

	//В чем полустрогость
	var uid = 1234
	fmt.Println("My UID", uid)
	fmt.Printf("%T\n", uid)

	//Декларирование и инициализация нескольких переменных
	var firstVar, secondVar int = 10, 20
	fmt.Printf("First: %d Second: %d\n", firstVar, secondVar)

	//Декларирование блока переменных
	var (
		personName string = "Bob"
		personAge  int    = 42
		personUID  int
	)
	fmt.Printf("Name: %s\nAge: %d\nUID: %d\n", personName, personAge, personUID)

	//Короткая декларация
	count := 1.2345
	fmt.Printf("%.2f\n", count)

	//Ввод
	var (
		myName string
		myAge  int
	)

	fmt.Scan(&myName, &myAge)

	fmt.Printf("My name is: %s\nAge is: %d\n", myName, myAge)

	//Ручное использование потока ввода
	fmt.Fscan(os.Stdin, &myAge)
	fmt.Println("New age", myAge)
}
