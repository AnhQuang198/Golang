package main

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width, height float64) Rectangle {
	return Rectangle{
		width:  width,
		height: height,
	}
}

func (r Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}
