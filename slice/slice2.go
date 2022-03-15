package main

import "fmt"

func main() {
	var s []int

	s = make([]int, 3)

	s[0] = 100
	s[1] = 200
	s[2] = 300

	t := append(s, 400)

	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))

	var u []int
	u = append(t, 500)
	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))
	fmt.Println(u, len(u), cap(u))

	u[0] = 9999

	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))
	fmt.Println(u, len(u), cap(u))
}
