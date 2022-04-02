package main

import "fmt"

func main() {
	// 아무것도 입력 안할 시 에러
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		panic(err)
	}

	fmt.Println(input)
}
