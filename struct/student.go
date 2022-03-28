package main

import "fmt"

type student struct {
	name   string
	gender string
	score  map[string]int
}

func newStudent(name, gender string) *student {
	return &student{name, gender, map[string]int{}}
}

func main() {
	var stdNum, classNum, score int
	var name, gender, className string

	fmt.Scanln(&stdNum, &classNum)

	s1 := make([]*student, 0, 0)

	for i := 0; i < stdNum; i++ {
		fmt.Scanln(&name, &gender)
		s1 = append(s1, newStudent(name, gender))

		for j := 0; j < classNum; j++ {
			fmt.Scanln(&className, &score)
			s1[i].score[className] = score
		}
	}

	for i := 0; i < stdNum; i++ {
		fmt.Println("----------")
		fmt.Println(s1[i].name, s1[i].gender)

		for index, val := range s1[i].score {
			fmt.Println(index, val)
		}

	}
	fmt.Println("----------")
}
