package main

import (
	"fmt"
	"strings"
)

func main() {
	// Условный оператор if
	var color string
	fmt.Println("Enter Color")
	fmt.Scan(&color)
	if strings.Compare("red", color) == 0 {
		fmt.Println("Color is Red")
	} else if strings.Compare("green", color) == 0 {
		fmt.Println("Color is green")
	} else {
		fmt.Println("Color Unknown")
	}

	// Инициализация в блоке условного оператора
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	//if data, err:= someFunc(); err != nil {
	//
	//}

}
