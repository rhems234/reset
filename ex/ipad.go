package main

import "fmt"

func inputNums() []int {
	var input int
	nums := make([]int, 5, 5)

	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &input)
		nums[i] = input
	}
	return nums
}

func calExam(arr []int) (sum int, over int, under int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] > 90 {
			over++
		} else if arr[i] < 70 {
			under++
		}
		sum += arr[i]
	}

	return
}

func printResult(sum int, over int, under int) {
	var result bool = true

	if sum < 400 {
		fmt.Println("총점이 400점 미만입니다.")
		result = false
	}
	if over < 2 {
		fmt.Println("90점 이상 과목 수가 2개 미만입니다.")
		result = false
	}
	if under > 0 {
		fmt.Println("70점 미만 과목이 있습니다.")
		result = false
	}

	if result == true {
		fmt.Println("아이패드를 살 수 있습니다.")
	} else {
		fmt.Println("아이패드를 살 수 없습니다.")
	}
}

func main() {
	nums := inputNums()
	printResult(calExam(nums))
}
