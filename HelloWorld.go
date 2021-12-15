package main

import "fmt"

var N = "Hellooooooo!!"
var str int

func main() {
	n := 1000
	i := "Hello"
	fmt.Println("Hello World!!!")
	fmt.Printf("%v, %T", i, i)
	fmt.Println(n)
	fmt.Println("Nhap so a: ")
	var a int
	fmt.Scan(&a)
	other(&a)
}

func other(a *int) {
	if *a%2 == 0 {
		fmt.Println(a, "la so chan")
	} else {
		fmt.Println(a, "la so le")
	}
}
