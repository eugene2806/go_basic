package main

import "fmt"

// 1. DEFER Оператор отложенного вызова функции/метода

func CheckDBCloseConnection(a int) {
	fmt.Println("Check started.... NumIp", a)
	fmt.Println("Check done! Connection refused")
}

// 2. Если функция принимает входные параметры и используется в связке с defer то:
// Параметры расчитываются в момент передачи их в функцию
// А вызов функции с уже давно расчитанными параметрами осущ. в момент завершения вышележащей функции

func main() {
	defer fmt.Println("Step 1") //Вызываются в порядке стека
	defer fmt.Println("Step 2")
	defer fmt.Println("Step 3")
	defer fmt.Println("Step 4")

	var numIP = 10
	p := &numIP
	defer CheckDBCloseConnection(numIP)
	*p++
	fmt.Println("Value numIP in main()", numIP)
	fmt.Println("Main function started")
	fmt.Println("Main function ended")
}
