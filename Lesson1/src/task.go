package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b int
	)
	fmt.Scan(&a, &b)

	if a+b < 100 && a > 0 && b > 0 {
		fmt.Println("Сумма чисел", a+b)
	} else {
		fmt.Println("Error")
	}

}
