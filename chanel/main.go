package main

import "fmt"

func main() {
	//unbuffered chanel
	fmt.Println("Start")
	ch := make(chan int)
	go func() {
		ch <- 100
		fmt.Println("Sent")
	}()
	fmt.Println("Process")
	fmt.Println(<-ch)
	fmt.Println("Done")

	fmt.Println("===============")

	//buffered chanel
	chBuf := make(chan int, 2) //type, cap

	chBuf <- 1
	fmt.Println("Process sent 1")
	chBuf <- 2
	fmt.Println("Process sent 2")

}
