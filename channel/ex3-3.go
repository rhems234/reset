package main

import "fmt"

func sum(num1, num2 int) <-chan int {
	result := make(chan int)

	go func() {
		result <- num1 + num2
	}()

	return result
}

func main() {
	ch := sum(10, 5)

	fmt.Println(<-ch)
}
