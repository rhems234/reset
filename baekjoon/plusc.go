package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, cycle int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var new = n

	for true {
		cycle++
		var sum int
		if new/10 == 0 {
			sum = new
		} else {
			sum = new/10 + new%10
		}
		new = new%10*10 + sum%10
		if new == n {
			break
		}
	}
	fmt.Println(cycle)
}
