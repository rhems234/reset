package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	input := make([]int, n)
	min := -1000000
	max := 1000000

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &input[i])

		if min < input[i] {
			min = input[i]
		} else if max > input[i] {
			max = input[i]
		}
	}
	fmt.Println(max, min)
}
