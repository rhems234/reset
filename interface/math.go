package main

import (
	"fmt"
	"math"
)

type project interface {
	vol() float64  // 부피
	area() float64 // 겉넓이
}

type Ball struct {
	radius, high float64
}

type Circle struct {
	a, b, c float64
}

func (cy Ball) vol() float64 {
	return (math.Pi * cy.radius * cy.radius) * cy.high
}

func (cy Ball) area() float64 {
	return (math.Pi*cy.radius*cy.radius)*2 + (2*math.Pi*cy.radius)*cy.high
}

func (cu Circle) vol() float64 {
	return cu.a * cu.b * cu.c
}

func (cu Circle) area() float64 {
	return 2*(cu.a*cu.b) + 2*(cu.a*cu.c) + 2*(cu.b*cu.c)
}

func main() {
	cy1 := Ball{10, 10}
	cy2 := Ball{4.2, 15.6}
	cu1 := Circle{10.5, 20.2, 20}
	cu2 := Circle{4, 10, 23}

	printMeasure(cy1, cy2, cu1, cu2)
}

func printMeasure(m ...project) {
	for _, val := range m {
		fmt.Printf("%.2f, %.2f\n", val.area(), val.vol())

	}
}
