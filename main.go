package main

import (
	"newproject/people"
)

func main() {
	students := []people.Student{}

	students = people.AddStudent(students, "Onur", "Yılmaz")
	students = people.AddStudent(students, "Onur", "Yılmaz")

	directory := people.ConvertToMap(students)

	people.PrintStudents(students)

	people.FindStudentById(directory, 1)

}
