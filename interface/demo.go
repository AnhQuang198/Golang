package main

import "fmt"

func main() {
	var animal Animal
	animal = Dog{}
	animal.Speak()
}

type Animal interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Speak: Go Go")
}
