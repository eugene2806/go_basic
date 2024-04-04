package main

import (
	"fmt"
	"math"
)

// 3. В чем преимущество методов над функциями?
// Наличие методов улучшает понимание кода
// В Go запрещается создавать функции с одинаковыми названиями, а медоды для разных структур
// С одинаковыми названиями разрешены

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, length int
}

// 4. Создадим функцию и  метод Perimeter для структуры Circle

func Perimeter(c Circle) float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) Perimeter() float64 {
	return c.radius * 2 * math.Pi
}

//5. Если создать функцию с названием Perimeter для Rectangle то ничего работать не будет

//func Perimeter(r Rectangle) float64 {
//	return c.radius * 2 * math.Pi
//}

// В то время как метод будет работать т.к принимает другую структуру

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func PerimeterRectangle(r Rectangle) int {
	return 2 * (r.length + r.width)
}

func main() {
	c := Circle{10.5}
	fmt.Println("Call function", Perimeter(c))
	fmt.Println("Call method", c.Perimeter())

	r := Rectangle{10, 5}
	fmt.Println("Call function", PerimeterRectangle(r))
	fmt.Println("Call method", r.Perimeter())
}
