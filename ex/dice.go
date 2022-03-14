package main

import "fmt"

func main() {
	var a int
	var b int
	var c int

	fmt.Scanf("%d %d %d\n", &a, &b, &c)

	if a == b && b == c {
		fmt.Printf("%d\n", (10000 + (a * 1000)))
	} else if a == b {
		fmt.Printf("%d\n", 1000+a*100)
	} else if b == c {
		fmt.Printf("%d\n", 1000+b*100)
	} else if c == a {
		fmt.Printf("%d\n", 1000+c*100)
	} else if a > b && a > c {
		fmt.Printf("%d\n", a*100)
	} else if b > a && b > c {
		fmt.Printf("%d\n", b*100)
	} else if c > a && c > b {
		fmt.Printf("%d\n", c*100)
	}
}
