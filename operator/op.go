package main

import "fmt"

func main() {
	a := 21
	c := a % 10
	a = a / 10
	d := a % 10

	fmt.Printf("첫번째 수 : %v, 두번쨰 수 : %v\n", c, d)
	/*	a := 4
		b := 2
		fmt.Printf("%v\n", a&b)
		fmt.Printf("%v\n", a|b)
		fmt.Println("result2=", a^b)
	*/
}
