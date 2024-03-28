package main

import "fmt"

// 1. Явные функции (Любая функция в Go) является экземпляром 1-го уровня
// Функцию можно присваивать в переменную, можно передавать в качестве параметра
// и передавать в другие функции

// 2. Возврат функции в качестве значения
func calcAndReturnValidFunc(command string) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }
	} else if command == "substraction" {
		return func(a, b int) int { return a - b }
	} else {
		return func(a, b int) int { return a * b }
	}
}

// 3. Функция как параметр в другой функции
func receiveFuncAndReturnValue(f func(a, b int) int) int {
	var intVarA, intVarB int
	intVarA = 100
	intVarB = 200
	return f(intVarA, intVarB)
}

func add(a, b int) int {
	return a + b
}

func main() {
	command := "addition"
	res := calcAndReturnValidFunc(command)
	fmt.Println("Result with command :", command, "value:", res(10, 20))
	//4. Тип функции

	fmt.Printf("Type of func => %T\n", res)
	//5. Тип функции в Go определяется как входными параметрами так и выходными

	fmt.Println("receiveFuncAndReturnValue(add): ", receiveFuncAndReturnValue(add))
	fmt.Println(receiveFuncAndReturnValue(func(a, b int) int {
		return a*a + b*b
	}))

}
