package main

import (
	"fmt"
	"reflect"
)

func main() {
	//khai bao khoi tao slice
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}

	c := append(a, b...)
	//for index, value := range a {
	//	fmt.Println(index, value)
	//}
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(a))

	//cap and len
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	e := d[0:2]
	f := d[:3]
	fmt.Printf("d %v, %v, %v\n", d, len(d), cap(d))
	fmt.Printf("e %v, %v, %v\n", e, len(e), cap(e))
	fmt.Printf("f %v, %v, %v\n", f, len(f), cap(f))

	g := append(f, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	fmt.Printf("g %v, %v, %v\n", g, len(g), cap(g))

	h := []int{}
	fmt.Printf("h %v, %v, %v\n", h, len(h), cap(h))

	m := make([]int, 0)
	n := append(m, 1, 2, 3)
	fmt.Printf("m %v, %v, %v\n", m, len(m), cap(m))
	fmt.Printf("n %v, %v, %v\n", n, len(n), cap(n))

	o := [3]int{1, 2, 3}
	q := o[:]
	p := append(h, q...)
	fmt.Printf("p %v, %v, %v\n", p, len(p), cap(p))
}
