package main

import "fmt"

func main() {
	c := make(chan string, 1)

	c <- "Hello goorm!"

	fmt.Println(<-c)
}
