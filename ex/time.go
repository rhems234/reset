package main

import "fmt"

func main() {
	h := 0
	m := 0
	var plus int
	g := 0
	fmt.Scanf("%d %d\n", &h, &m)
	fmt.Scanf("%d", &plus)
	m += plus
	if m >= 60 {
		g = m / 60
		m = m - 60*g
	}
	h = h + g
	if h >= 24 {
		h = h - 24
	}
	fmt.Printf("%d %d", h, m)
}
