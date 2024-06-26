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
				*NewLessonInfo(lesson, grade),
			},
		}

		return
	}

	addLesson(studentCard, student, lesson, grade, checkLesson)

}

func NewLessonInfo(newLesson string, newGrade int) *LessonInfo {
	return &LessonInfo{Title: newLesson, Grade: newGrade}
}

func createLesson(
	studentCard map[string]Student,
	student Student,
	lesson string,
	grade int,
) {
	student.Lessons = append(student.Lessons, *NewLessonInfo(lesson, grade))
	studentCard[student.Name] = student
}

func addLesson(
	studentCard map[string]Student,
	student Student,
	lesson string,
	grade int,
	checkLesson bool,
) {
	if checkLesson {
		for i, les := range student.Lessons {
			if les.Title == lesson {
				student.Lessons[i].Grade = grade

				return
			}
		}
	}

	createLesson(studentCard, student, lesson, grade)

	return
}

func (s Student) averageGrade() (float32, bool) {
	var sum int
	var res float32
	var counter int

	if s.Name == "" {
		fmt.Printf("Ошибка подсчета для: %s", s.Name)

		return 0, false
	}

	for _, lesson := range s.Lessons {
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
		student, _ := studentCard[studentName].averageGrade()
		fmt.Println("Средняя оценка по всем предметам:", student)
	}
}

func checkLessonDuplicate() bool {
	var answer string
	fmt.Println("Исключать дубли предметов из карты студентов?")

	for {
		fmt.Printf("Введите ответ YES | NO : ")

		_, err := fmt.Scan(&answer)
		if err != nil {
			fmt.Println("Ошибка ввода")

			continue
		}

		answer = strings.ToLower(answer)

		if answer == trueQuestionAnswer {

			fmt.Println("Дубли предметов будут удалены")

			return true
		}

		if answer == falseQuestionAnswer {

			fmt.Println("Дубли предметов будут добавлены")

			return false
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
	}

	checkLesson := checkLessonDuplicate()

	for _, student := range studentArray {
		addStudentCard(studentGrades, student[0], student[1], student[2], checkLesson)
	}

	fmt.Println("================")

	studentName1, success1 := studentGrades["Anna"].averageGrade()
	fmt.Printf("Средняя оценка студента Anna %.2f. Подсчет верный: %t\n", studentName1, success1)
	studentName2, success2 := studentGrades["Mike"].averageGrade()
	fmt.Printf("Средняя оценка студента Mike %.2f. Подсчет верный: %t\n", studentName2, success2)

	fmt.Println("================")

	studentList(studentGrades)

}
