package main

import "fmt"

var result chan int = make(chan int)

func add(num1 *int, num2 *int) {
	result <- *num1 + *num2
}

func main() {
	var num1, num2 int

	fmt.Scanln(&num1, &num2)

	go add(&num1, &num2)

	result := <-result
	fmt.Println(result)
}
