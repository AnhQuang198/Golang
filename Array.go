package main

import (
	"fmt"
)

func main() {
	var myArray [5]int
	fmt.Println(myArray)
	fmt.Println(len(myArray))
	array1 := [...]int{5, 6, 7}
	fmt.Println(array1)

	//loop array
	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	for index, value := range array1 {
		fmt.Printf("i=%d value=%d", index, value)
	}

	for _, value := range array1 {
		fmt.Println(value)
	}

	//mang 2 chieu [row][column]
	matrix := [4][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}
	fmt.Println(matrix)

	//loop matrix
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(matrix[i][j])
		}
		fmt.Println()
	}
}
