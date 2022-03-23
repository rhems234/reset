package main

import "fmt"

func printSqure(a *int) {
	*a *= *a

	fmt.Println(*a)
}

func main() {
	a := 4
	printSqure(&a)

	fmt.Println(a)
}
