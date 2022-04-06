package main

import "fmt"

func main() {
	c := make(chan string, 1) // 버퍼

	c <- "Hello goorm!"

	fmt.Println(<-c)
}
