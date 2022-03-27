package main

import "fmt"

const g float32 = 9.8

type Energy func(float32, float32) float32

func calMechEnergy(f Energy, a float32, b float32) float32 {
	result := f(a, b)
	return result
}

func main() {
	var m, v, h float32
	fmt.Scanln(&m, &v, &h)

	kinEnergy := func(m float32, v float32) float32 {
		return (m * (v * v) / 2)
	}
	potEnergy := func(m float32, h float32) float32 {
		return m * g * h
	}

	ke := kinEnergy(m, v)
	pe := potEnergy(m, h)
	fmt.Printf("%v %v %v", ke, pe, pe+ke)
}
