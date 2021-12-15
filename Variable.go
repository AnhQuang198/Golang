package main

import (
	"fmt"
	"math"
)

func main() {
	//Bool
	var isCheck bool = true
	fmt.Print(isCheck)

	//Int
	var num int = 123
	fmt.Println(num)

	//Int 8, 16, 32, 64
	//1. Range
	//2. Bits

	//range int8 = -128 -> 127
	//max 8 bit
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)

	//range int16 = -32768 -> 32767
	//max 16 bit
	fmt.Println(math.MinInt16)
	fmt.Println(math.MaxInt16)

	//range int32 = -2147483648 -> 2147483647
	//max 32 bit
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)

	//range int64 = -9223372036854775808 -> 9223372036854775807
	//max 64 bit
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)

	var mUint uint = 10 //so nguyen duong
	fmt.Println(mUint)

	var mByte byte = 1 //giong kieu uint8
	fmt.Println(mByte)
	fmt.Println("%T", mByte) //uint8

	var a byte = 'A'
	fmt.Println(a) //65 -> index A of ASCII

	//Float - so thuc
	var myFloat float32 = 23.4
	fmt.Println(myFloat)

	//Complex - so phuc
	var myComplex complex64 = 10 + 1i //phan thuc + phan ao
	fmt.Println(myComplex)

	//Rune - khong quan tam den so bit ma bien tao ra
	//type conversion
	var y int = 1
	var z float32 = 1.2
	//fmt.Println(y + z) error -> khong the tu ep kieu trong go
	fmt.Println(float32(y) + z) //true

	//const
	const PI = 3.14 //character, string, bool, numeric
	//PI := 3.14 -> error
	fmt.Println(PI)
}
