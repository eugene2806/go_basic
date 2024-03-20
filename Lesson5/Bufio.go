package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var name string
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() { //Команда захвата потока ввода и сохранения в буфер (Захват идет до символа окончания строки)
		name = input.Text() //Команда возвращения элементов помещенных в буфер (Отдаст string)
	}
	fmt.Println(name)

	fmt.Println("For loop started:")
	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println("Input string is ", result) // Выводит текст пока не получит на вход пустую строку
		}
	}

	//Преобразование строкового литерала к чему-нибудь числовому
	numStr := "10"
	numInt, _ := strconv.Atoi(numStr)
	fmt.Printf("%d is %T\n", numInt, numInt)
}
