package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, c float64
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &a, &b, &c)

	if c-b <= 0 {
		fmt.Println(-1)
		return
	}
	round := int(math.Ceil(a / (c - b)))

	if round == int(a/(c-b)) {
		fmt.Println(round + 1)
	} else {
		fmt.Println(round)
	}

}
