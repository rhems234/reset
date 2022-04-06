package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a byte
	reader := bufio.NewReader(os.Stdin)
	a, _ = reader.ReadByte()

	fmt.Printf("%d\n", a)
}
