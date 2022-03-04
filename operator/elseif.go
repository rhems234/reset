package main

import "fmt"

func main() {
	a := 5
	if a < 10 && a > 2 /*a == 3 || a == 4 */ {
		fmt.Println("a는 10보다 작고 2보다 크다")
	} else {
		fmt.Println("a는 10보다 크거나 a는 2보다 작다")
	}
	fmt.Println("a의 값은 ", a)
	/*a := 5
	if a == 3 {
		fmt.Println("a 는 3")
	} else if a == 4 {
		fmt.Println("a 는 4")
	} else {
		fmt.Println("a는 3도 아니고 4도 아니다.")
	}
	fmt.Println("a의 값은 ", a) */
}
