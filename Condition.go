package main

import "fmt"

func main() {
	//if else
	var a int = 10
	//basic
	if a > 5 {
		fmt.Print("So lon hon 5")
	} else {
		fmt.Println("So nho hon 5")
	}

	//advanced
	//if statement; condition {}
	if z := 10; z > 5 {
		fmt.Println("So lon hon 5")
	} else {
		fmt.Println("So nho hon 5")
	}

	//switch case
	number := 10
	switch {
	case number == 10:
		fmt.Println("Day la 10")
		fallthrough //tiep tuc chay ma ko break case
	case number == 20:
		fmt.Println("Day khong phai la 10")
	}
}
