package main

import "fmt"

// 3*2없애기
func main() {
	var dan int
	var j int

	if dan == 3 && j == 2 {
		for dan = 1; dan < 10; dan++ {
			fmt.Printf("%d단\n", dan)

			for j = 1; j < 10; j++ {
				fmt.Printf("%d * %d = %d\n", dan, j, dan*j)
			}
			continue
		}
		fmt.Println()
	}
}
