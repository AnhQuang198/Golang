package main

import "fmt"

func main() {
	//declare and set value pointer
	a := 100.1
	var pointer *float64

	pointer = &a
	fmt.Println(pointer)
	fmt.Printf("%T", pointer)

	var point *int
	fmt.Println()
	fmt.Println(point)

	pointer2 := new(int)
	fmt.Println(pointer2)

	//change value pointer
	num := 10
	var numPoint *int
	numPoint = &num

	fmt.Println("======")
	fmt.Println(num)
	//set value pointer num
	*numPoint = 2000

	fmt.Println(numPoint)
	fmt.Println(num)

	//pointer -> array
	array := [3]int{1, 2, 3}
	var arrPoint *[3]int
	arrPoint = &array
	fmt.Println(arrPoint)
	fmt.Printf("%T", arrPoint)

	fmt.Println("========")
	fmt.Println(num)
	getpointerValue(&num)
}

func getpointerValue(pointer *int) {
	*pointer = 1000
	fmt.Println("Poinnnnter", *pointer)
}
