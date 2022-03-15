package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var numbers = make([]int, 9)
	reader := bufio.NewReader(os.Stdin)
	var max = 0
	var maxindex = 0

	for i := range numbers {
		fmt.Fscanf(reader, "%d\n", &numbers[i])
		if numbers[i] > max {
			max = numbers[i]
			maxindex = i
		}
	}
	fmt.Println(max)
	fmt.Println(maxindex + 1)
}
