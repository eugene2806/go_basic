package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func reverse(arr []int) []int {
	var newSlice []int
	for i := len(arr) - 1; i >= 0; i-- {
		newSlice = append(newSlice, arr[i])
	}
	return newSlice
}

func convertCsvFile(s []string) []int {
	var numbers  []int
	for _, value := range s {
		num, _ := strconv.Atoi(value)
		numbers = append(numbers, num)
	}
	return numbers
}

func readCsvFile() []int {
	file, err := os.Open("/Users/eugene/Desktop/go/go_basic/src/input.csv")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return nil
	}
	defer file.Close()
	reader := csv.NewReader(file)
	record, e := reader.Read()
	if e != nil {
		fmt.Println("Ошибка чтения файла:", e)
		return nil
	}
	fmt.Println(record)
	slice := convertCsvFile(record)
	return slice
}

func writeCsvFile(nums []int) {
	file, err := os.Create("/Users/eugene/Desktop/go/go_basic/src/output.csv")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	var numsString []string
	for _, val := range nums {
		numsString = append(numsString, fmt.Sprintf("%d", val))
	}
	err = writer.Write(numsString)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}
	defer writer.Flush()
	fmt.Println("Файл успешно записан")
}

func main() {
	slice := readCsvFile()
	res := reverse(slice)
	writeCsvFile(res)
	fmt.Println(res)
}
