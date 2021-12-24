package main

import "fmt"

func main() {
	hello()
	sayHello("Quang")
	result := nice("quang")
	fmt.Println(result)

	width, height := recInfo(5, 6)
	fmt.Print("Width: ", width)
	fmt.Println("Height: ", height)

	w, h := recInfo1(6, 7)
	fmt.Println("W: ", w)
	fmt.Println("H: ", h)
	//Call()
}

//func khong truyen vao param va khong co gia tri tra ve
func hello() {
	fmt.Println("Hello")
}

//func co tham so truyen vao
func sayHello(name string) {
	fmt.Println("Hello", name)
}

//func có tham số truyền vào và có trả về
func nice(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

//multiples return value
func recInfo(width, height int) (int, int) {
	return width, height
}

//Named return value
func recInfo1(width, height int) (w, h int) {
	w = width
	h = height
	return w, h
}
