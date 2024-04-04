package main

import "fmt"

//1. Если интерфейс пустой то ему может удовлетворять кто угодно

type Empty interface {
}

//2. Реализуем функцию которая может удовлетворять кому угодно

func Describer(pretendent interface{}) {
	fmt.Printf("Type %T and Value %v\n", pretendent, pretendent)
}

type Student struct {
	name string
}

//4.Type Assertion

func SemiGeneric(pretendent interface{}) {
	val, ok := pretendent.(int)
	fmt.Println("Value", val, "Ok?", ok)
}

func DoSomething(pretendent interface{}) {
	switch pretendent.(type) { //Пытаемся извлечь нижележащий тип
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	default:
		fmt.Println("Unknown Type")
	}
}

func main() {
	strSample := "eecrecec"
	//3. Передача параметра без присваивания в промежуточную переменную
	Describer(strSample)

	intSample := 200
	Describer(intSample)

	Describer(Student{name: "Alex"})
	fmt.Println()

	//5. Работа с полу-дженериками
	SemiGeneric(10)
	SemiGeneric(Student{name: "Alex"})
	SemiGeneric("dsdsd")

	DoSomething(10)
	DoSomething("abcd")
	DoSomething(func(a, b int) int { return a + b })
}
