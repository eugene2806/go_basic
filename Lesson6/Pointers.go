package main

import "fmt"

// 1. Указатели - переменная, хранящая в качестве значения - адрес памяти другой переменной

func main() {
	//2. Определение указателя на что-то
	var variable int = 30
	var pointer *int = &variable //&... Операция взятия адреса в памяти
	// В pointer лежит 1803434xcd0000132 - это место в памяти где хранится int значение 30

	fmt.Printf("Type of pointer : %T\n", pointer)  // *int
	fmt.Printf("Value if pointer : %v\n", pointer) // 0x1400000e088

	//3. Нулевоке значение для указателя - nil
	var zeroPointer *int
	fmt.Printf("ZeroPointer : %v\n", zeroPointer)

	//4. Разыменование указателя - получение значения| *pointer возвращает значение хранимое по адресу
	var numericValue int = 32
	var pointerToNumeric *int = &numericValue
	fmt.Printf("Value numeric: %v\nAddress is %v\n", *pointerToNumeric, pointerToNumeric)
	//& - Применяется для переменных чтобы создать указатель
	//* - Применяется для указателя чтобы получить значение

	//5. Создание указателей на нулевые значения типов
	//var zeroVar int
	//var zeroPoint *int = &zeroVar
	zeroPoint := new(int) // Создает под капотом zeroValue для int, и возвращает адресс где этот 0 хранится
	fmt.Printf("Value in *zeroPoint = %v\nAddress is %v\n", *zeroPoint, zeroPoint)

	//6. Изменение значения хранимого по адресу через указатель
	zeroPointerToInt := new(int)
	fmt.Printf("Address is %v and Value zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt)
	*zeroPointerToInt += 40
	fmt.Printf("Address is %v and New Value zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt)

	//7. Указательная арифметика ОТСУТСТВУЕТ ПОЛНОСТЬЮ (Нельзя прибавить ссылку к ссылке)

	//8. Передача указателей в функции
	//Прирост производительности за счет передачи указателя а не копии значения
	sample := 1
	//samplePointer := &sample
	fmt.Println("Origin value sample", sample)
	//changeParam(samplePointer)
	changeParam(&sample) //Прямое создание указателя
	fmt.Println("After change sample", sample)

	//9. Возвратпоинтера из функции
	ptr1 := returnPointer()
	ptr2 := returnPointer()
	fmt.Printf("Ptr1: %T and address %v and value %v\n", ptr1, ptr1, *ptr1)
	fmt.Printf("Ptr1: %T and address %v and value %v\n", ptr2, ptr2, *ptr2)

	//10. Передача слайса на массив в функцию
	newArr := [3]int{1, 2, 3}
	fmt.Println("Before mutation", newArr)
	mutationSlice(newArr[:])
	fmt.Println("After mutation", newArr)

}

// 8.1 Определение функции принимающей параметр как указатель
func changeParam(val *int) {
	*val += 100
}

// 9.1 Определение функции, возвращающей указатель
func returnPointer() *int {
	var numeric int = 321
	return &numeric //В момент возврата Go перемещает данную переменную в кучу
}

// 10.1 Функция мутации массива
func mutationSlice(sls []int) {
	sls[1] = 100
	sls[2] = 200
}
