package main

import "fmt"

func main() {
	var m = make(map[string]int)
	var avg float32
	var total int
	var score int
	var sub string

	for {
		fmt.Scan(&sub)
		if sub == "0" {
			break
		}
		fmt.Scan(&score)
		m[sub] = score
	}

	for index, num := range m {
		if index == "0" {
			break
		}
		total = total + num
		fmt.Printf("%s %d\n", index, num)
	}

	avg = float32(total) / float32(len(m))
	fmt.Printf("%.2f", avg)
}
