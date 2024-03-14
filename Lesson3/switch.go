package main

import "fmt"

func main() {
	//Switch
	var price int
	fmt.Scan(&price)

	switch price {
	case 100:
		fmt.Println("First Price")
	case 200:
		fmt.Println("Second Price")
	default:
		fmt.Println("default price")
	}

	//Case с множеством вариантов
	var ageGroup string = "b" // Возрастные группы: "a", "b", "c", "d", "e"
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("AgeGroup 10-40")
	case "d", "e":
		fmt.Println("AgeGroup 50-70")
	default:
		fmt.Println("You are too yong/old")
	}

	//Case с выражениями
	var age int = 17

	switch {
	case age <= 18:
		fmt.Println("You are yang")
	case age > 18 && age <= 30:
		fmt.Println("To old")
	default:
		fmt.Println("Who are you")
	}

	//Case с проваливаниями. Проваливания выполняют даже ложные кейсы
	//В момент выполнения fallthrough у следующего кейса не проверяется условие
	switch {
	case age <= 18:
		fmt.Println("You are yang")
		fallthrough
	case age > 18 && age <= 30:
		fmt.Println("To old")
	default:
		fmt.Println("Who are you")
	}

	//Териминация цикла for из switch
	var i int
uberloop:
	for {
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Println("Value is Even")
			break uberloop
		}
	}
}
