package main

import "fmt"

func guide() { // 매개변수, 반환 값 x
	fmt.Println("두 정수 입력 : ")
}

func input() (int, int) { // 매개변수 x, 반환값 o
	var a, b int
	fmt.Scanln(&a, &b)
	return a, b
}

func multi(a, b int) int { // 매개변수, 반환 값 o
	return a * b
}

func PrintResult(num int) { // 매개변수 o, 반환 값 x
	fmt.Printf("결과 값 :  %d\n", num)
}

func main() {
	guide()
	num1, num2 := input()
	result := multi(num1, num2)
	PrintResult(result)
}
