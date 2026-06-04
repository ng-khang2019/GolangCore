package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	name  string
	id    string
	class string
}

func inputStudent(list []Student, numberStudents int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---Input student information---")

	for i := 0; i < numberStudents; i++ {
		fmt.Println("Student", i+1, "information")

		fmt.Print("Name: ")
		name, _ := reader.ReadString('\n')
		list[i].name = strings.TrimSpace(name)

		fmt.Print("ID: ")
		fmt.Scanln(&list[i].id)
		fmt.Print("Class: ")
		fmt.Scanln(&list[i].class)
	}
}

func printStudent(list []Student) {
	fmt.Println("---List of all students:---")
	for _, student := range list {
		fmt.Printf("Name: %s, ID: %s, Class: %s\n", student.name, student.id, student.class)
	}
}

func main() {

	fmt.Print("Input the number of students: ")
	var numberStudents int
	fmt.Scan(&numberStudents)
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	students := make([]Student, numberStudents)
	inputStudent(students, numberStudents)
	fmt.Println()
	printStudent(students)
}
