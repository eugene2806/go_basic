package main

import "fmt"

//Почему интерфейс полу-абстрактный тип в Go

// 1. Создадим интерфейс ездилка

type Rider interface {
	Ride()
	Gas()
	Stop()
}

func main() {
	// 2. Создадим экземпляр ездилки
	var r Rider //ZeroValue = nil, ZeroType = nil
	if r == nil {
		fmt.Printf("r is nil. Value %v and Type %T\n", r, r)
	}

	r.Ride() //Тут будет паника т.к мы можем вызвать метод Ride(), но значение которое там лежит
	//это nil получается nil.Ride(). Поэтому экземпляру интерфейса должен быть присвоен претендент
	//реализующий методы
}
