package main

import (
	"fmt"
	"unicode/utf8"
)

// 1. Описание того что должен уметь претендент

type BigWord interface {
	IsBig() bool
}

// 2. Претендент которого нужно научить делать IsBig

type MySuperString string

// 3. Реализация IsBig() у MySuperString

func (mss MySuperString) IsBig() bool {
	if utf8.RuneCountInString(string(mss)) > 10 {
		return true
	}
	return false
}

func main() {
	sample := MySuperString("sjdjur")
	var interfaceSample BigWord //Объявление переменной, типа интерфейса BigWord
	interfaceSample = sample    //Присваивание значения для переменной типа BigWord возможно
	//потому что sample (типа MySuperString) удовлетворяет интерфейсу BigWord

	fmt.Println("IsBig?", interfaceSample.IsBig())

	//interfaceSample = "adfadf" - не будет работать так как тип string не имеет реализации IsBig
	//Поэтому не удовлетворяет интерфейсу BigWord

}
