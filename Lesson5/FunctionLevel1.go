package main

import "fmt"

func main() {
	rectanglePerimeter, rectangleArea := rectangleParameters(10, 10)
	fmt.Println("Perimeter, area: ", rectanglePerimeter, rectangleArea)

	namedPerimeter, namedArea := namedReturn(5, 5)
	fmt.Println("NamedPerimeter, NamedArea: ", namedPerimeter, namedArea)

	helloVariadic(10, 20, 30, 40)
	helloVariadic()

	// 4.Анонимная функция
	anon := func(a, b int) int {
		return a + b
	}

	fmt.Println("Anon", anon(3, 4))

	fmt.Println("bigFunction", bigFunction(10, 20))

}

// Функции
// 1.Возврат больше чем одного значения
func rectangleParameters(length, width float64) (float64, float64) {
	var perimeter = 2 * (length + width)
	var area = length * width
	return perimeter, area
}

// 2.Именованый возврат значений
func namedReturn(a, b int) (perimeter, area int) {
	perimeter = 2 * (a + b)
	area = a * b
	return //Не нужно указывать возвращаемые значения
}

// 3.Variadic parameters (Принимает несколько значений преобразуя их в слайс)
func helloVariadic(a ...int) { //Континуальные параметры указываются в конце и может быть только один
	fmt.Printf("Value %v and type %T\n", a, a)
	//Если нужно передать в функцию слайс
	//slice:= []int{1, 2, 3}
	//helloVariadic(slice...)
}

// 5.Анонимная функция внутри явной
func bigFunction(aArg, bArg int) int {
	return func(a, b int) int { return a + b }(aArg, bArg)
}
