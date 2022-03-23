package main

import "fmt"

var a int = 1 // 전역변수 선언

func localVar() int { // 지역변수로 연산
	var a int = 10
	a += 3
	return a
}

func globalVar() int { // 전역변수로 연산
	a += 3

	return a
}

func main() {
	fmt.Println("지역변수 a : ", localVar())
	fmt.Println("전역변수 a : ", globalVar())
}
