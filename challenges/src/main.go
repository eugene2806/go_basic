package main

import "fmt"

type studentMap map[string]map[string]int

func addStudentCard(studentCard studentMap, name, lesson string, grade int) {
	var answer string
	if lessonMap, ok := studentCard[name]; ok {
		if value, isOk := lessonMap[lesson]; isOk {
			fmt.Printf("Вы хотите изменить данные по уже существующему предмету %s %s!\n ", name, lesson)

			for {
				fmt.Printf("Вы уверены что хотите внести изменения? YES | NO :")
				_, err := fmt.Scan(&answer)
				if err == nil && (answer == "YES" || answer == "yes") {
					lessonMap[lesson] = grade
					fmt.Println("Изменения внесены")
					return
				}
				if err == nil {
					if answer == "NO" || answer == "no" {
						lessonMap[lesson] = value
						fmt.Println("Изменения не внесены")
						return
					}
					fmt.Println("Ошибка ввода")
				}
			}
		}
		lessonMap[lesson] = grade
		return
	}
	studentCard[name] = map[string]int{lesson: grade}
}

func averageGrade(studentCard studentMap, name string) (float32, bool) {
	var sum int
	var res float32
	var counter int
	if lessonMap, ok := studentCard[name]; ok {
		for _, value := range lessonMap {
			sum += value
			counter++
		}
		res = float32(sum) / float32(counter)
		fmt.Println("==============>", name)
		return res, true
	}
	fmt.Printf("Ошибка подсчета для: %s", name)
	return 0, false
}

func studentList(studentCard studentMap) {
	for studentName, lesson := range studentCard {
		fmt.Printf("Имя слудента: %s\n", studentName)
		for currentLesson, grade := range lesson {
			fmt.Printf("Предмет=> %s| Оценка: %d\n", currentLesson, grade)
		}
	}
}
func main() {
	studentGrades := make(studentMap)
	addStudentCard(studentGrades, "John", "Математика", 8)
	addStudentCard(studentGrades, "Anna", "Русский", 5)
	addStudentCard(studentGrades, "Anna", "ИЗО", 8)
	addStudentCard(studentGrades, "Mike", "Риторика", 6)
	addStudentCard(studentGrades, "Gleb", "Геометрия", 2)
	addStudentCard(studentGrades, "Gleb", "Изо", 3)
	addStudentCard(studentGrades, "Anna", "Английский", 7)
	// addStudentCard(studentGrades, "Anna", "Английский", 8)

	studentName1, success1 := averageGrade(studentGrades, "Anna")
	fmt.Printf("Средняя оценка студента %.2f. Подсчет верный: %t\n", studentName1, success1)
	studentName2, success2 := averageGrade(studentGrades, "Mike")
	fmt.Printf("Средняя оценка студента %.2f. Подсчет верный: %t\n", studentName2, success2)

	studentList(studentGrades)
}
