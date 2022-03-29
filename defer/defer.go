package main

import "fmt"

func main() {
	var a, b int = 10, 0
	defer fmt.Println("Done")

	result := a / b
	fmt.Println(result)
}
