package main

import (
	"fmt"
)

func validate() int {
	var (
		a int
		b int
	)
	fmt.Println("Введите первое число")
	_, err := fmt.Scan(&a)
	fmt.Println("Введите второе число")
	_, err2 := fmt.Scan(&b)
	if err != nil && err2 != nil {
		fmt.Println("Error")
		return 0
	}
	if a+b > 100 || a < 0 || b < 0 {
		fmt.Println("Error")
		return 0
	}
	return a + b
}
func main() {
	fmt.Println("Сумма чисел", validate())

}
