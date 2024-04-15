package main

import (
	"fmt"
	"modul2/circle"
	"modul2/rectangle"
)

// 1. Функция init данная функция отрабатывет единожды при первом импортировании пакета
// 2. Данных функций в пакете может быть несколько штук (Но не в одном модуле)

// 3. init() вызывается в момент инициализации пакета
// ** Сначала компилятор смотрит на содержимое пакета
// ** Затем компилятор смотрит на пути импорта (если что-то импортируется компилятор уходит туда)
// ** Затем копилятор инициализирует переменные уровня пакета
// ** Затем компилятор запускает функцию init() для данного пакета
// ** Повторяет данную процедуру для всех пакетов проекта
// ** После чего вызывается функция main()

//4. Что произойдет, если запустить go run main.go
// * Сначала смотрим в main.go на предмет синтаксических ошибок , но ничего не инициализируется
// * Затем импорты : сначала импоритруем Lec23/rectangle
// ** Компилятор идет в rectangle
// ** Смотрим в пакет на предмет синтаксических ошибок
// ** Затем импорт fmt
// ** Затем инициализируем переменные уровня пакета
// ** Затем запускаем функцию init() пакета rectangle
// ** Затем подружаем (определяем) все имена пакета rectangle
// ** Функции main() тут нет, возвращаемся назад
// * Пытаемся импоритровать fmt (т.к. он уже был импортирован одним из пакетов - повторный импорт не требуется)
// * Инициализируем переменные уровня пакета main
// * Запускаем функцию init() в main
// * Затем определяем имена (тут дополнительных имен нет, тут ничего не делаем)
// * Затем запускаем функцию main()

// 5. Все импорты (вне зависимости стандартные или пользовательские) сортируются по алфавиту

func init() {
	fmt.Println("Init function from main package!")
}

func main() {
	r := rectangle.New(10, 10, "red")
	fmt.Println(r)
	c := circle.Circle{Radius: 22.5}
	fmt.Println(c)
}
