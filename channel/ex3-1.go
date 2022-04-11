package main

import "fmt"

func channel1(ch chan int) {
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("done1")
}

func channel2(ch chan int) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch <- 3
	ch <- 4

	fmt.Println("done2")
}

func main() {
	c := make(chan int)

	go channel1(c)
	go channel2(c)

	fmt.Scanln()
}
