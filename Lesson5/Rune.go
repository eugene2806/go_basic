package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Hello World"
	fmt.Println(name)
	//1.Строка это байтовый слайс со своими особенностями при обращении
	//к нижележащему массиву
	word := "Тестовая строка"
	fmt.Printf("String: %s\n", word)
	//Каие значения байс сейчас находятся в слайсе word?
	fmt.Printf("Bytes:")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i]) //%x - фотмат представления 16-ти ричного байта
	}
	fmt.Println()
	//Получаем доступ к отдельно стоящим символам
	fmt.Printf("Characters: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i]) //%с - фотмат представления символа
	}
	fmt.Println()
	//2.Строки в Go - хранятся в виде UTF-8 символов.
	//При использовании кирилицы символы занимают не 1 байт, а 2 байта

	//3. Руна это стандартный встроенный тип в Go(alias над int32) позволяющий хранить
	//единый и неделимый UTF символ ВНЕ Зависисти от того сколько байт он занимает
	runeSlice := []rune(word) //Преобразование слайса байт к слайсу рун []byte(rune)
	fmt.Printf("runeSlice: ")
	//for i := 0; i < len(runeSlice); i++ {
	//	fmt.Printf("%c ", runeSlice[i])
	//}
	for _, char := range runeSlice {
		fmt.Printf("%c ", char)
	}
	fmt.Println()
	//4.Создание строки из слайса байт
	myByteSlice := []byte{0x40, 0x41, 0x42, 0x43}
	myStr := string(myByteSlice)
	fmt.Println("myByteSlice:", myStr)

	//5.Руны как литералы
	runeLiteralSlice := []rune{'V', 'a', 's', 'y', 'a'}
	myStrFromRuneLiterals := string(runeLiteralSlice)
	fmt.Println("myStrFromRuneLiterals", myStrFromRuneLiterals)

	//6.Длинна и емкость строки
	//Длинна len()- количество байт в слайсе
	fmt.Println("Length of Вася: ", len("Вася"), "bytes")
	//Длинна   RuneCounter - количество рун!
	fmt.Println("Length of Вася: ", utf8.RuneCountInString("Вася"), "runes")
	//Вычисление емкости строки бессмысленно т.к строка базовый тип

	//7. Сравнение строк == и != начиная с версии Go 1.6
	word1, word2 := "Vasya", "Petya"
	if word1 == word2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
	//8. Конкатенация
	word3 := word1 + " " + word2
	fmt.Println(word3)

	//9.Строитель строк
	firstName := "Alex"
	secondName := "Johnson"
	fullName := fmt.Sprintf("%s ### %s", firstName, secondName) //Возвращает форматированную строку
	fmt.Println(fullName)

	//10.Строки не изменяемые
	//fullName[0] = 'Q'
	//11. А слайсы изменяемые
	fullNameSlice := []rune(fullName)
	fullNameSlice[0] = 'F'
	firstName = string(fullNameSlice)
	fmt.Println("String Mutation: ", firstName)

	//12. Методы для работы со строками
	//Пакет import "strings"
}
