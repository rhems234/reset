package main

import "fmt"

func main() {
	var i int
	for {
		if i == 5 {
			i++
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
		i++
	}

	/*var i int
	for {
		i++
		fmt.Println(i)
	}

	/*for i = 0; i < 10; i++ {
		fmt.Println(i)
	}
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++ // i = i + 1
	}*/
}
