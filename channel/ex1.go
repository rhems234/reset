package main

import "fmt"

func main() {
	var a, b = 10, 5
	var result int

	go func() {
		result = a + b
	}()

	fmt.Printf("두 수의 합은 %d 입니다.", result)
}
