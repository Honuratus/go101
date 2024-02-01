package people

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	ID        int
}

func PrintStudents(students []Student) {
	for _, student := range students {
		fmt.Println(student)
	}
}

func FindStudentById(sMap map[int]string, id int) {
	name, exists := sMap[id]
	if !exists {
		fmt.Printf("id: %d found\n", id) 
		return
	}
	fmt.Printf("[%d]%s found\n", id, name) 
}

func AddStudent(students []Student, firstName string, lastName string) []Student {
	student := Student{
		FirstName: firstName,
		LastName:  lastName,
		ID:        len(students) + 1,
	}
	students = append(students, student)
	return students
}

func ConvertToMap(students []Student) map[int]string {
	resultMap := make(map[int]string)

	for _, student := range students {
		resultMap[student.ID] = student.FirstName + " " + student.LastName
	}

	return resultMap
}
