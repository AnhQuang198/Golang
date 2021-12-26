package main

import (
	"HelloWorld/oop/department"
	employee "HelloWorld/oop/employee"
	"fmt"
)

func main() {
	department := department.New("IT Innovation")
	department.ShowInfo()

	fmt.Println()

	e := employee.Employee{
		FirstName: "Le Anh",
		LastName:  "Quang",
	}
	e.GetInfo()
}
