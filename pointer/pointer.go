package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintSungjuk() {
	fmt.Println(s.class, s.grade, s.name, s.age)
}

func (s *Student) InputSungjuk(class string, grade string, name string, age int) {
	s.class = class
	s.grade = grade
	s.name = name
	s.age = age
}

func main() {
	var s Student = Student{name: "JunSeong", age: 24, class: "수학", grade: "A+"}
	s.InputSungjuk("박준성", "과학", "C", 24)
	s.PrintSungjuk()
}
