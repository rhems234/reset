package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &a)

	var count = 1
	var room = 1

	for true {
		if count >= a {
			break
		}
		count += 6 * room
		room++
	}

	fmt.Println(room)
}
