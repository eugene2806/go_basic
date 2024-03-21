package main

import "fmt"

func main() {
	//1.Map - это набор пар ключ:значение. Инициализация пустой мапы
	mapper := make(map[string]int)
	fmt.Println("mapper", mapper)

	//2.Добавление пар в существующую мапу
	mapper["Alice"] = 24
	mapper["Bob"] = 25
	fmt.Println("Mapper after adding pairs: ", mapper)

	//3.Инициализация мапы с указанием пар
	newMapper := map[string]int{
		"Alice": 1000,
		"Bob":   1500,
	}
	newMapper["Jo"] = 3000
	fmt.Println("newMapper", newMapper)

	//4.Что может быть ключом в мапе?
	//4.1 В Go мапа не упорядочена
	//4.2 Ключом в мапе может быть только сравнимый тип (==, !=)

	//5.Нулевое значение для map == nil
	//var mapZeroValue map[string]int

	//6.Получение элементов из мапы
	//6.1 Получение элемента, который представлен в мапе
	testPerson := "Alice"
	fmt.Println("Salary of test person: ", newMapper[testPerson])
	//6.2 Получение элемента который Не представлен в мапе
	testPerson = "Derek" // При обращении к несуществующему ключу вернется zeroValue для заданного типа
	fmt.Println("Salary of test person: ", newMapper[testPerson])

	//7.Проверка вхождения ключа
	employee := map[string]int{
		"Den":   0,
		"Alice": 0,
		"Bob":   0,
	}
	fmt.Println("Den:", employee["Den"])
	fmt.Println("Jo:", employee["Jo"])
	//7.1 Прои обращении по ключу возвращается 2 значения
	if value, ok := employee["Jo"]; ok { //Значение и true/false
		fmt.Println("Jo and value", value)
	} else {
		fmt.Println("Jo does not exist in map")
	}

	//8.Перебор элементов мапы
	fmt.Println("==========")
	for key, value := range employee {
		fmt.Printf("Key - %s and value - %d\n", key, value)
	}

	//9.Удаление пар delete(employee, "Den")
	//9.1 Удаление не существующей пары
	if _, ok := employee["Anna"]; ok {
		delete(employee, "Anna") //Дорогая операция, поэтому нужна проверка по ключу
	}
	fmt.Println("After deleting: ", employee)

	//10.Количество пар == длинна мапы
	fmt.Println("Pair amount in map:", len(employee))

	//11. Мапа как и слайс ссылочный тип (При копировании будет мутировать исходник)

	//12.Сравнение мап
	//12.1 Сравнение массивов (массив можно использовать как ключ в мапе)
	if [3]int{1, 2, 3} == [3]int{1, 2, 3} {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
	//12.2 Слайсы нельзя сравнивать
	//12.2 Мапы как и слайсы можно сравнивать только с == nil
	//Пустая не проинициализированная мапа == nil
}
