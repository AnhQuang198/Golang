package main

type Square struct {
	width  float64
	height float64
}

func NewSquare(width, height float64) Square {
	square := Square{
		width:  width,
		height: height,
	}
	return square
}

func (s Square) Perimeter() float64 {
	return s.height * 4
}
