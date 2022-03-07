package main

import "fmt"

func main() {
	s := 0
	for i := 1; i <= 10; i++ {
		s += i
	}
	fmt.Println(s)

	/*s := sum(10, 0)
	fmt.Println(s) */
}

/*func sum(x int, s int) int {
	if x == 0 {
		return s
	}
	s += x
	return sum(x-1, s)
} */

/*func f1(x int) {
	if x == 0 {
		return
	}
	fmt.Printf("f1(%d) before call f1(%d)\n", x, x-1)
	f1(x - 1)
	fmt.Printf("f1(%d) after call f1(%d)\n", x, x-1)
}*/
