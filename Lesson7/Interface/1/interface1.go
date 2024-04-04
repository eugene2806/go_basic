package main

import "fmt"

// 1. Интерфейсы - явно декларированный набор сигнатур ПОВЕДЕНИЙ
//Набор методов удовлетворив которые, можно считаться типом который декларирует интерфейс
//Если структура описывает паттерн состояния то интерфейс описывает Паттерн Поведения!
// Интерфейсы в Go декларируют ПОЛУ-АБСТРАКТНЫЙ тип

type FigureBuilder interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	length, width int
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

type Circle struct {
	radius int
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}

func (c Circle) Area() int {
	return 3 * c.radius * c.radius
}

func main() {
	// 2. Создадим по 3 экземпляра
	r1 := Rectangle{10, 15}
	r2 := Rectangle{15, 20}
	r3 := Rectangle{5, 25}
	c1 := Circle{5}
	c2 := Circle{10}
	c3 := Circle{15}

	//3. Посчитаем общую площадь этих фигур
	rectangles := []Rectangle{r1, r2, r3}
	totalSumAreaOfRectangles := 0
	for _, rect := range rectangles {
		totalSumAreaOfRectangles += rect.Area()
	}

	circles := []Circle{c1, c2, c3}
	totalSumAreaOfCircles := 0
	for _, circ := range circles {
		totalSumAreaOfCircles += circ.Area()
	}

	fmt.Println("Total Area is:", totalSumAreaOfCircles+totalSumAreaOfRectangles)

	//4. Реализуем подсчет через интерфейс
	figures := []FigureBuilder{r1, r2, r3, c1, c2, c3}
	total := 0
	for _, figure := range figures {
		total += figure.Area()
	}

	fmt.Println("Total: ", total)
}
