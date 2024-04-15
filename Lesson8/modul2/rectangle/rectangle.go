package rectangle

import "fmt"

func init() {
	fmt.Println("Init function from rectangle package!")
	fmt.Println("Name:", name, "Age:", age)
}

// 1. Инициализируем переменные уровня пакета

var (
	name string = "John"
	age  int    = 33
)

type Rectangle struct {
	A, B  int
	color string
}

func New(newA, newB int, newColor string) *Rectangle {
	return &Rectangle{newA, newB, newColor}
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.A + r.B)
}
