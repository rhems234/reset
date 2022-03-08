package main

import "fmt"

func main() {
	s := "Hello World"

	s2 := []rune(s)
	fmt.Println("len(s2) = ", len(s2))
	for i := 0; i < len(s2); i++ {
		fmt.Print(s2[i], ", ")
	}
	/*fmt.Println("len(s) = ", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]), ",")
	}*/
}
