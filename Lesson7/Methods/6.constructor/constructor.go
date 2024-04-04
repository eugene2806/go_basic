package main

import "fmt"

//1. Создадим структуру

type Rectangle struct {
	length, width int
}

// 2. Создадим конструкто для Rectangle
// Для имен конструкторов используют следующие наименования:
// * New() - если данный конструктор один в файле (В файле присутствует описание только одной структуры)
// * New<StructName>() - если в файле присутствуют еще структуры

// 3. В Go принято возвращать из конструктора не сам экземпляр а ссылку на него

func New(newLength, newWidth int) *Rectangle {
	return &Rectangle{newLength, newWidth}
}

// 4. Добавим 2 метода

func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	r := New(10, 15)
	fmt.Printf("r %v\n", r) // r являтся указателем (r &{10 15})
	fmt.Println("Perimeter", r.Perimeter())
	fmt.Println("Area", r.Area())
}
