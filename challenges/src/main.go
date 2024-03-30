package main

import (
	"fmt"
	"strconv"
	"strings"
)

const trueQuestionAnswer = "yes"
const falseQuestionAnswer = "no"

type Student struct {
	Name    string
	Lessons []LessonInfo
}

type LessonInfo struct {
	Title string
	Grade int
}

// Добавление студента
func addStudentCard(
	studentCard map[string]Student,
	name, lesson, gradeStr string,
	checkLesson bool,
) {

	grade, _ := strconv.Atoi(gradeStr)
	student, ok := studentCard[name]

	if !ok {
		studentCard[name] = Student{
			Name: name,
			Lessons: []LessonInfo{
				{lesson, grade},
			},
		}
		return
	}

	if checkLesson == false {
		student.Lessons = append(student.Lessons, LessonInfo{Title: lesson, Grade: grade})
		studentCard[name] = student

		return
	}
	checkLessonDuplicate(studentCard, student, name, lesson, grade)
}

func checkLessonDuplicate(
	studentCard map[string]Student,
	student Student,
	name, lesson string,
	grade int,
) {

	var answer string

	for i, les := range student.Lessons {
		if les.Title == lesson {
			fmt.Printf(
				"Вы хотите изменить данные по уже существующему предмету %s %s на %d\n ", name, lesson, grade,
			)

			for {
				fmt.Printf("Вы уверены что хотите внести изменения? YES | NO :")

				_, err := fmt.Scan(&answer)
				if err != nil {
					fmt.Println("Ошибка ввода")

					return
				}

				if strings.ToLower(answer) == trueQuestionAnswer {
					student.Lessons[i].Grade = grade

					fmt.Println("Изменения внесены")

					return
				}

				if strings.ToLower(answer) == falseQuestionAnswer {
					fmt.Println("Изменения не внесены")

					return
				}
			}
		}
	}
	student.Lessons = append(student.Lessons, LessonInfo{Title: lesson, Grade: grade})
	studentCard[name] = student
}

func averageGrade(studentCard map[string]Student, name string) (float32, bool) {
	var sum int
	var res float32
	var counter int

	student, ok := studentCard[name]

	if !ok {
		fmt.Printf("Ошибка подсчета для: %s", name)

		return 0, false
	}

	for _, lesson := range student.Lessons {
		sum += lesson.Grade
		counter++
	}
	res = float32(sum) / float32(counter)

	return res, true
}

func studentList(studentCard map[string]Student) {
	for studentName, lesson := range studentCard {
		fmt.Printf("Имя слудента: %s\n", studentName)

		for _, currentLesson := range lesson.Lessons {
			fmt.Printf("Предмет=> %s| Оценка: %d\n", currentLesson.Title, currentLesson.Grade)
		}
	}
}

func main() {
	studentGrades := make(map[string]Student)
	studentArray := [][]string{
		{"John", "Математика", "8"},
		{"Anna", "Русский", "5"},
		{"Anna", "ИЗО", "8"},
		{"Mike", "Риторика", "6"},
		{"Gleb", "Геометрия", "2"},
		{"Gleb", "Изо", "3"},
		{"Anna", "Английский", "7"},
		{"Anna", "Английский", "8"},
		{"Gleb", "Геометрия", "5"},
	}

	for _, student := range studentArray {
		addStudentCard(studentGrades, student[0], student[1], student[2], true)
	}

	fmt.Println("================")
	studentName1, success1 := averageGrade(studentGrades, "Anna")
	fmt.Printf("Средняя оценка студента Anna %.2f. Подсчет верный: %t\n", studentName1, success1)
	studentName2, success2 := averageGrade(studentGrades, "Mike")
	fmt.Printf("Средняя оценка студента Mike %.2f. Подсчет верный: %t\n", studentName2, success2)

	fmt.Println("================")
	studentList(studentGrades)

}
