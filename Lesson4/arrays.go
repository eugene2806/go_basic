package main

import "fmt"

func main() {
	//Массивы. Основа
	//1.Определение массива
	var arr [5]int //При инициализации массиива указываем сколько элементов в нем будет
	fmt.Println(arr)
	//Обращение к элементу массива arr[i] = elem
	arr[0] = 100
	arr[1] = 106
	arr[2] = 6
	fmt.Println("Add elements arr:", arr)
	//2.Определение массива с указанием элементов
	newArr := [5]int{100, 150, 12, 21, 43}
	fmt.Println("newArr:", newArr)
	//4.Создание массива через инициализацию переменных
	arrWithValues := [...]int{1, 2, 4, 94, 11}
	fmt.Println("arrWithValues:", arrWithValues)
	//5.Массив это набор значений. Т.е при всех манипуляциях - массив копируется(Жестко на уровне компилятора)
	first := [...]int{1, 2, 3}
	second := first
	second[0] = 10000
	fmt.Println("First:", first)
	fmt.Println("Second:", second)
	//7.Многомерный массив
	words := [2][2]string{
		{"Bob", "Alice"},
		{"Alex", "Zoi"},
	}
	fmt.Println("Words", words)
	//8.Итерирование по многомерному массиву
	for _, val1 := range words {
		for _, val2 := range val1 {
			fmt.Printf("%s ", val2)
		}
		fmt.Println()
	}
	
}
