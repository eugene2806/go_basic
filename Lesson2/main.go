package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//Numerics. Integers
	//uint8, uint16, uint32, uint64
	//int8, int16, int32, int64
	var a int = 32
	b := 90
	//1.
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))
	//При использовании короткого объявления тип для целого числа -int платформозависимый
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b))

	//2. Явно приводить типы нужно от меньшего к большему чтобы не потерять значения в случае int32(second64)
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(int64(first32) + second64)

	//Numerics. Float
	//float32, float64
	floatFirst, floatSecond := 3.56, 7.64
	fmt.Printf("floatFirst %.3f, floatSecond %.3f\n", floatFirst, floatSecond)

	//Numerics. Complex
	c1 := complex(5, 7)
	fmt.Println(c1)
	c2 := 12 + 32i
	fmt.Println(c2)

	//Strings. Строка это набор БАЙТ
	name := "Федя"
	lastName := "Pupkin"
	concat := name + " " + lastName
	fmt.Println("Full name:", concat)
	fmt.Println("Length of string:", name, len(name))
	fmt.Println("Amount of chars:", name, utf8.RuneCountInString(name))

	//Rune - это один utf-ный символ.
	//Поиск подстроки в строке
	totalString, subString := "ABCDERT", "CDE"
	fmt.Println(strings.Contains(totalString, subString))
	//rune -> alias int32
	sampleRune := 'Q' // Для инициализыции руны символьным значением - используем ''
	anotherRune := 1048
	fmt.Printf("Rune as char: %c\n", sampleRune)
	fmt.Printf("Rune as char: %c\n", anotherRune)
	// "A" < "abcd"
	fmt.Println(strings.Compare("abc", "abcd")) // -1 of a<b, 0 if a==b, 1 a>b

	var aByte byte // byte-> int8
	fmt.Println(aByte)
}
