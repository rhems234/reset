package main

import "fmt"

// 3*2없애기
func main() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i == 3 && j == 2 {
				continue
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
}
