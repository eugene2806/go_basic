package main

import "fmt"

// 1. Структура это заименованный набор полей (свойств), определяющий новый тип данных.

// 2. Определение структуры (явное определение)

type Student struct {
	firstName string
	lastName  string
	age       int
}

// 3. Если имеется ряд состояний одного типа, можно сделать так

type AnotherStudent struct {
	firstName, lastName, groupName string
	age, courseNumber              int
}

// 9. Структуры с анонимными полями
type Human struct {
	firstName string
	string
	int
	bool
}

func printStudent(stud Student) {
	fmt.Println("========")
	fmt.Println("firstName:", stud.firstName)
	fmt.Println("lastName:", stud.lastName)
	fmt.Println("age:", stud.age)
}

func main() {
	//4. Создание представителей структуры
	stud1 := Student{
		firstName: "Alex",   //Наименования полей можно не прописывать
		lastName:  "Petrov", //Но тогда должен быть строгий порядок наименований полей
		age:       22,
	}
	printStudent(stud1)

	// 5. Анонимные структуры (структура без имени)
	anonStudent := struct {
		age           int
		groupID       int
		professorName string
	}{
		23,
		2,
		"Alekseev",
	}
	fmt.Println("anonStudent", anonStudent)

	// 6. Доступ к состояниям
	studAlex := Student{"Alex", "NeAlex", 22}
	fmt.Println("studAlex age:", studAlex.age)
	studAlex.age += 2
	fmt.Println("studAlex age:", studAlex.age)

	// 7. Инициализация пустой структуры
	emptyStudent := Student{}
	var emptyStudent2 Student
	printStudent(emptyStudent)
	printStudent(emptyStudent2)

	// 8. Указатели на экземпляры структур
	studPointer := &Student{
		"Igor",
		"Sidorov",
		30,
	}
	fmt.Println("studPointer", studPointer)
	secondPointer := studPointer
	secondPointer.age += 4
	//(*secondPointer).age += 4 //Можно явно не делать (*...) разыменование указателя т.к это делает сам компилятор
	//studPointer.age += 4
	fmt.Println("studPointer", studPointer)
	studPointerNew := new(Student) //Возвращает ссылку на zeroValue структуры
	fmt.Println("studPointerNew", studPointerNew)

	// 9.1 Создание экземпляра структур с анонимными полями
	human := &Human{
		"Bob",
		"Jonson",
		22,
		true,
	}
	fmt.Println("human", human)
	fmt.Println("human.int", human.int)
	
}
