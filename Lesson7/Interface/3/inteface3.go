package main

import "fmt"

type Worker interface {
	Work()
}

type Employee struct {
	name string
	age  int
}

type Student struct {
	name         string
	courseNumber int
}

func (s Student) Work() {
	fmt.Println("Now Student with name", s.name, "is working!")
}

func (e Employee) Work() {
	fmt.Println("Now Employee with name", e.name, "is working!")
}

func Describer(w Worker) {
	fmt.Printf("Interface is type %T and value %v\n", w, w)
}
func main() {

	emp := Employee{"Bob", 33}
	var workerEmployee Worker = emp //Присваиваем сотрудника в переменную типа Worker
	workerEmployee.Work()
	Describer(workerEmployee) //В результате видим, что тип интерфейса определяется структурой
	//его реализующей, а значение интерфейса это значение состояний структуры

	std := Student{"Alex", 19}
	var workerStudent Worker = std
	workerStudent.Work()
	Describer(workerStudent)

	var workers = []Worker{workerStudent, workerEmployee}
	for _, worker := range workers {
		Describer(worker)
	}
}
