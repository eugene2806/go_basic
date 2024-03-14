package main

import "fmt"

func main() {
	//Вложенные циклы и лейблы
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//Лейблы для остановки цикла
outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d j:%d and sum i+j:%d\n", i, j, i+j)
			if i == j {
				break outer // Завершает работу всех циклов
			}
		}
	}

	//Классический цикл While
	var loopVar int = 0
	for loopVar <= 5 {
		fmt.Println(loopVar)
		loopVar++
	}

	// Классический бесконечный цикл
	var password string
outer2:
	for {
		fmt.Println("Enter password")
		fmt.Scan(&password)
		if password != "12345" {
			fmt.Println("Try again")
		} else {
			fmt.Println("Password accepted")
			break outer2
		}
	}

}
