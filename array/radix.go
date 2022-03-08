package main

import "fmt"

func main() {
	// radix sort 특정 사용 조건
	// 원소 값의 범위가 한정적일때,
	// 알고리즘 속도는 배열의 갯수만큼 필요
	arr := [11]int{0, 5, 4, 9, 1, 2, 8, 3, 6, 4, 5}
	temp := [10]int{}

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}
	fmt.Println(arr)
}
