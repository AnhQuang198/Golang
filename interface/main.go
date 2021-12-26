package main

import "fmt"

func main() {
	var shape Shape
	shape = NewRectangle(1.5, 5.5)
	perRectangle := shape.Perimeter()
	fmt.Printf("Dien tich hinh chu nhat: %v", perRectangle)

	var shapeSquare Shape = NewSquare(1.5, 5.5)
	perSquare := shapeSquare.Perimeter()
	fmt.Printf("Dien tich hinh vuong: %v", perSquare)
}
