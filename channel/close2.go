package main

import "fmt"

func main() {
	c := make(chan string, 2)

	c <- "Hello"
	c <- "goorm"

	close(c)

	val, open := <-c
	fmt.Println(val, open)
	val, open = <-c
	fmt.Println(val, open)
	val, open = <-c
	fmt.Println(val, open)
	val, open = <-c
	fmt.Println(val, open)
}
