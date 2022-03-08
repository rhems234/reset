package main

import "fmt"

type Student struct {
	name  string
	class int

	sungjuk Sungjuk
}

type Sungjuk struct {
	name  string
	grade string
}

func (s Student) ViewSungJuk() {
	fmt.Println(s.sungjuk)
}

func (s Student) InputSungjuk(name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

func InputSungjuk(s Student, name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

func ViewSungJuk(s Student) {
	fmt.Println(s.sungjuk)
}
func main() {
	var s Student
	s.name = "철수"
	s.class = 1

	s.sungjuk.name = "수학"
	s.sungjuk.grade = "C"

	s.ViewSungJuk()
	ViewSungJuk(s)

	s.InputSungjuk("과학", "A")
	s.ViewSungJuk()
}
