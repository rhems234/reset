package main

import "fmt"

const g float32 = 9.8

type gravity struct {
	m, v, h    float32
	ke, pe, me float32
}

func (ke gravity) kegy() float32 {
	return ke.m * (ke.v * ke.v) / 2
}

func (pe gravity) pegy() float32 {
	return pe.m * g * pe.h
}

func main() {
	var l, s, high float32
	var num int

	fmt.Scanln(&num)
	g1 := make([]gravity, num)

	for i := 0; i < num; i++ {
		fmt.Scanln(&l, &s, &high)
		g1[i] = gravity{}
		g1[i].m = l
		g1[i].v = s
		g1[i].h = high
		g1[i].ke = g1[i].kegy()
		g1[i].pe = g1[i].pegy()
		g1[i].me = g1[i].ke + g1[i].pe
	}

	for i := 0; i < num; i++ {
		fmt.Println(g1[i].ke, g1[i].pe, g1[i].ke+g1[i].pe)
	}
}
