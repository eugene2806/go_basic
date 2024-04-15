package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strconv"
)

func funcWithError(a int) (string, error) {
	if a%2 == 0 {
		return "Even", nil
	}

	return "", errors.New("THIS IS ERROR!")
}

func PanicRecover() {
	if r := recover(); r != nil {
		fmt.Println("Panic satisfied", r)
		debug.PrintStack()
	}
}
func panicExample(firstName *string, lastName *string) {
	defer PanicRecover()
	if firstName == nil {
		panic("Runtime error: firstName nil")
	}
	if lastName == nil {
		panic("Runtime error: lastName nil")
	}
	fmt.Println("Full name", firstName, lastName)
}

func main() {

	numLiteral := "12"
	num, err := strconv.Atoi(numLiteral)
	if err != nil {
		fmt.Println("Can not convert this value to int:", err)
		return
	}
	fmt.Println("Convertion done:", num)

	name := "Bob"
	panicExample(&name, nil)

	ans, err := funcWithError(5)
	if err != nil {
		fmt.Println("Not use odd ", err)
		return
	}
	fmt.Println(ans)

}
