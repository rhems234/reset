package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input int
	var max, sum int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &input)

	var score = make([]int, input)
	for i := 0; i < input; i++ {
		fmt.Fscanf(reader, "%d", &score[i])
		sum += score[i]
		if max < score[i] {
			max = score[i]
		}
	}
	fmt.Println(float64(sum) / float64(input) / float64(max) * 100.0)
}
