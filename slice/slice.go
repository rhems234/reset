package main

import "fmt"

func main() {
	a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2

	b := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}

	b = append(b, 3)

	fmt.Printf("%p %p\n", a, b)

	b[0] = 4
	b[1] = 5

	fmt.Println(a, b)
	/*a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2
	b := append(a, 3)

	fmt.Printf("%p %p\n", a, b)
	fmt.Println(a)
	fmt.Println(b)

	b[0] = 4
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)*/

	/*a := []int{1, 2}

	b := append(a, 3)

	fmt.Printf("%p %p\n", a, b)

	for i := 0; i < len(a); i++ {
		fmt.Printf("%d", a[i])
	}
	fmt.Println()

	for i := 0; i < len(b); i++ {
		fmt.Printf("%d", b[i])
	}
	fmt.Println()

	fmt.Println(cap(a), " ", cap(b))*/

}
