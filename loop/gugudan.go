package main

import "fmt"

func main() {

	for dan := 1; dan <= 9; dan++ {
		/*if dan == 5 {
			continue
		}*/ //5단 빼기
		fmt.Printf("%d단\n", dan)

		for j := 1; j <= 9; j++ {
			/*if j == 5 {
				continue
			}*/ //각 단의 5만 뺴기
			fmt.Printf("%d * %d = %d\n", dan, j, dan*j)
		}

		fmt.Println()
	}
}
