package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d\n", &a)
	fmt.Fscanf(reader, "%d", &b)

	fmt.Println(a + b)
}
