package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) PrintName() {
	fmt.Print(p.name)
}

func main() {
	var p Person
	p1 := Person{"Kevin", 15}
	p2 := Person{name: "Jon", age: 21}
	p3 := Person{name: "Carson"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)

	p.name = "Smith"
	p.age = 24
	fmt.Println(p)

	p.PrintName()
}
