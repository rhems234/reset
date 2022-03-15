package main

import "fmt"

func main() {
	var n int
	var x int
	var input int

	fmt.Scanf("%d %d\n", &n, &x)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &input)

		if x > input {
			fmt.Printf("%d", input)
		}
	}
}
