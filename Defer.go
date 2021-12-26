package main

import "fmt"

func main() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Error ", error)
		}
	}()

	a := 10
	b := 1
	//catch exception div by zero
	fmt.Println(a / b)
	fmt.Println("End")
}
