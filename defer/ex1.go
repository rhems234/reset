package main

import "fmt"

func main() {
	names := make([]string, 0, 3)
	var name string

	for {
		fmt.Scanln(&name)
		if name == "0" {
			break
		} else {
			names = append(names, name)
		}
	}

	for _, v := range names {
		defer fmt.Println(v)
	}
}
