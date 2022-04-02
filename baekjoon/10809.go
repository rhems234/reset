package main

import (
	"bufio"
	"encoding/ascii85"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &a)

	var b = make(map[int]int)
	for i := 0; i < 26; i++ {
		b[i+97] = -1
	}
	for i := 0; i < len(a); i++ {
		ascii, _ := strconv.Atoi(fmt.Sprint())
	}
}