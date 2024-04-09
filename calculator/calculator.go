package main

import (
	"fmt"
	"os"
	"strconv"
)

const fileName = "calculatorOutput.txt"

// сервис, который выводит информацию в терминал

type TerminalService struct {
}

type FileService struct{}

func (s *TerminalService) Save(text string) {
	fmt.Println("Ответ:", text)
}

func (sf *FileService) FileSave(text string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка при создании файла")

		return
	}

	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Ошибка при записи в файл")

		return
	}

	fmt.Println("Результат успешно записан в ", fileName)
}

// сервис, который что-то считает

type Calculator struct {
}

func (c *Calculator) calc(x, y int) int {
	return x + y
}

type MainService struct {
	saver      TerminalService
	fileSaver  FileService
	calculator Calculator
}

func (ms *MainService) run(x, y int) {
	result := ms.calculator.calc(x, y)  // Считаем через калькулятор
	ms.saver.Save(strconv.Itoa(result)) // Сохраняет результат
	ms.fileSaver.FileSave(strconv.Itoa(result))
}

func main() {
	mainService := MainService{
		saver:      TerminalService{},
		fileSaver:  FileService{},
		calculator: Calculator{},
	}

	mainService.run(10, 20)
}

// main.go -> MainService -> Calculator && TerminalSaver
// main.go -> MainService -> Calculator && FileSaver
