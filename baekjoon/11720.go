package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	fmt.Fscanf(reader, "%s", &input)

	inputs := strings.Split(input, "")

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(inputs[i])
		m += num
	}
	fmt.Println(m)
}
