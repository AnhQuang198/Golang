package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
}

type Class struct {
	id          int
	teacherName string
	student     Student
}

func main() {
	//named declare struct (nen su dung)
	st1 := Student{
		id:   1,
		name: "Quang",
	}
	fmt.Println(st1)
	fmt.Println(st1.id)
	fmt.Println(st1.name)

	//declare fast
	st2 := Student{2, "Quang Anh"}
	fmt.Println(st2)

	//anonymous struct
	var ano = struct {
		email string
		age   int
	}{
		"leequang198@gmail.com",
		23,
	}
	fmt.Println(ano)

	fmt.Println("======")
	//pointer struct
	pointer := &Student{ //pointer tro den dia chi cua struct Student
		id:   2,
		name: "Anh Quang",
	}
	fmt.Println((*pointer).id)
	fmt.Println((*pointer).name)

	//struct long struct -> Nested Struct

}
