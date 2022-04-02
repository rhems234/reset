package main

import (
	"errors"
	"fmt"
)

func inputSubNum() (num int, err error) {
	fmt.Scanln(&num)

	if num == 0 {
		num = 0
		err = errors.New("잘못된 과목 수입니다.")
		return
	}

	err = nil
	return
}

func average(num int) (avg float64, err error) {
	var score, total int

	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		if score < 0 || score > 100 {
			avg = 0
			err = errors.New("잘못된 점수입니다.")
			return
		}
		total += score
	}

	avg = float64(total) / float64(num)

	err = nil
	return
}

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()

	num, err1 := inputSubNum()

	if err1 != nil {
		panic(err1)
	}

	result, err2 := average(num)

	if err2 != nil {
		panic(err2)
	}

	fmt.Println(result)
}
