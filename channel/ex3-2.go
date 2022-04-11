package main

import "fmt"

func sendChannel(ch chan<- int) {
	ch <- 1
	ch <- 2
	// <-ch 오류 발생
	fmt.Println("done1")
}

func receiveChannel(ch <-chan int) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//ch <- 3 오류 발생
	fmt.Println("done2")
}

func main() {
	c := make(chan int)

	go sendChannel(c)
	go receiveChannel(c)

	fmt.Scanln()
}
