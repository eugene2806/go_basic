package main

import (
	"fmt"
	"time"
)

func newGoRoutine() {
	fmt.Println("Hey, i'm new Goroutine")
}

// 2. Функция main() является горутиной
// Особенность в том что если main() завершается то все горутины тоже завершаются
func main() {
	go newGoRoutine() //В этот момент происходит формирование запроса на вызов функции в отдельной горутине.
	// соответственно код основной горутины main() продолжет выполняться сразу же и не ждет завершения
	// newGoRoutine()
	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine work!")
	//Запустим код и.....
}
