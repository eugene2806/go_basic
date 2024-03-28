package main

import "fmt"

// 1. Вложенные структуры (вложение структур) - это использование одной структуры как
// тип поля в другой структуре

type University struct {
	yearBased int
	infoShort string
	infoLong  string
}

type Student2 struct {
	firstName  string
	secondName string
	university University
}

// 4. Встроенные структуры (Когда мы добавляем поля одной структуры в другую)

type Professor struct {
	firstName  string
	secondName string
	University // В этом месте происходит добавление всех полей Uni в Professor
}

func main() {
	//2. Создание экземпляров структур с вложением
	stud := Student2{
		"Alex",
		"NeAlex",
		University{
			1991,
			"Cool Univer",
			"Very cool Univer",
		},
	}
	fmt.Println(stud)

	//3. Получение доступа к полям структур
	fmt.Println("stud.firstName", stud.firstName)
	fmt.Println("stud.university.yearBased", stud.university.yearBased)

	//5. Создание экземпляра с встраиванием структур
	prof := Professor{
		firstName:  "Bob",
		secondName: "NeBob",
		University: University{
			1860,
			"Long info",
			"Short info",
		},
	}
	fmt.Println(prof)

	//6. Обращение к состояниям с встроенной структурой (Реализация наследования)
	fmt.Println("Prof.firstName", prof.firstName)
	fmt.Println("Prof.infoShort", prof.infoShort) // Обращаемся к полям University напрямую

	//7. Сравнение экземпляров ==
	//При сравнении экземпляров происходит сравнение всех их полей
	profLeft := Professor{}
	profRight := Professor{}
	fmt.Println(profLeft == profRight)

	//8. Если хотя бы одно из полей не сравнимо, ТО ився структура не сравнима
	//К примеру Map, slice

}
