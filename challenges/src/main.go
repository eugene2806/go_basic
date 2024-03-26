package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addStudentCard(studentCard map[string]map[string]int,
	name, lesson string,
	gradeStr string) {

	grade, _ := strconv.Atoi(gradeStr)
	var trueQuestionAnswer string = "yes"
	var falseQuestionAnswer string = "no"
	var answer string

	lessonMap, ok := studentCard[name]
	value, isOk := lessonMap[lesson]

	if !ok {
		studentCard[name] = map[string]int{lesson: grade}

		return
	}

	if !isOk {
		lessonMap[lesson] = grade

		return
	}

	if isOk {
		fmt.Printf("Вы хотите изменить данные по уже существующему предмету %s %s!\n ", name, lesson)

		for {
			fmt.Printf("Вы уверены что хотите внести изменения? YES | NO :")

			_, err := fmt.Scan(&answer)
			if err != nil {
				fmt.Println("Ошибка ввода")

				return
			}

			if strings.ToLower(answer) == trueQuestionAnswer {
				lessonMap[lesson] = grade

				fmt.Println("Изменения внесены")

				return
			}

			if strings.ToLower(answer) == falseQuestionAnswer {
				lessonMap[lesson] = value

				fmt.Println("Изменения не внесены")

				return
			}
		}
	}
}

func averageGrade(studentCard map[string]map[string]int, name string) (float32, bool) {
	var sum int
	var res float32
	var counter int

	lessonMap, ok := studentCard[name]

	if !ok {
		fmt.Printf("Ошибка подсчета для: %s", name)

		return 0, false
	}

	if ok {
		for _, value := range lessonMap {
			sum += value
			counter++
		}
		res = float32(sum) / float32(counter)

		return res, true
	}

	return 0, false
}

func studentList(studentCard map[string]map[string]int) {
	for studentName, lesson := range studentCard {
		fmt.Printf("Имя слудента: %s\n", studentName)

		for currentLesson, grade := range lesson {
			fmt.Printf("Предмет=> %s| Оценка: %d\n", currentLesson, grade)
		}
	}
}

func main() {
	studentGrades := make(map[string]map[string]int)
	studentArray := [][]string{
		{"John", "Математика", "8"},
		{"Anna", "Русский", "5"},
		{"Anna", "ИЗО", "8"},
		{"Mike", "Риторика", "6"},
		{"Gleb", "Геометрия", "2"},
		{"Gleb", "Изо", "3"},
		{"Anna", "Английский", "7"},
		{"Anna", "Английский", "8"},
	}

	for _, student := range studentArray {
		addStudentCard(studentGrades, student[0], student[1], student[2])
	}

	fmt.Println("================")
	studentName1, success1 := averageGrade(studentGrades, "Anna")
	fmt.Printf("Средняя оценка студента Anna %.2f. Подсчет верный: %t\n", studentName1, success1)
	studentName2, success2 := averageGrade(studentGrades, "Mike")
	fmt.Printf("Средняя оценка студента Mike %.2f. Подсчет верный: %t\n", studentName2, success2)

	fmt.Println("================")
	studentList(studentGrades)

}
