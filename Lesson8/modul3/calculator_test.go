package main

import (
	"log"
	"testing"
)

// 1. Файл с модульными тестами принято называть:
// * <script_name>_test.go
// * <package_name>_test.go

// 2. Для того чтобы тестировать (методы, структуры, интерфейсы и т.д)
// На каждый юнит создаем по своей тестирующей функции (Test)
// Принято каждую функцию называть со слова Test...

func TestAdd(t *testing.T) {
	// 1-й test-case
	if res := Add(10, 20); res != 29 {
		log.Fatalf("Add (10, 20) fail. expected %d, got %d\n", 30, res)
		// log.Fatal провоцирует завершение работы
	}
	if res := Add(30, 30); res != 60 {
		log.Fatalf("Add (10, 20) fail. expected %d, got %d\n", 60, res)
		// log.Fatal провоцирует завершение работы
	}
}

func testMult(t *testing.T) {

}

// 3. Для запуска тестов используем команду go test
// Детально : go test -v
// go test -cover показывает процент покрытия тестами кода
// 70-80% coverage этого бывает более чем достаточно

// 4. Все что начинается со слова Test будет запущено командой go test
// В Go принято что создается 1 модуль с тестами на весь пакет (вне зависимости от количества модулей в нем)
// Не тестируйте Getattr/Setattr в общем пайплайне (только специфика)
// Обязательно посмотрите как происходит обвязка с CI для Go (TravisCI/CircleCI)
