package main

import "fmt"

func main() {
	a := 4
	if a == 3 {
		fmt.Println("a 는 3")
	} else {
		fmt.Println("a 는 3이 아니다")
	}
	fmt.Println("a의 값은 ", a)
	/*	var a bool
		a = 3 < 4 && 2 < 5
		fmt.Println(a) */
}
