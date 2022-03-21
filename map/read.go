package main

import "fmt"

func main() {
	// 지역번호와 지역 저장
	var m = make(map[string]string)

	m["02"] = "서울특별시"
	m["031"] = "경기도"
	m["032"] = "충청남도"
	m["053"] = "대구광역시"

	fmt.Println(m["032"])
	fmt.Println(m["042"], "빈 칸 입니다.")
	// string형태로 존재하지 않는 KEY값은 ""가 출력된다.

	val, exist := m["02"] // 존재하지 않는 key
	fmt.Println(val, exist)

	val, exist = m["042"]
	fmt.Println(val, exist) // 존재하지 않는 key

	val = m["053"]
	fmt.Println(val)

	_, exist = m["053"]
	fmt.Println(exist)

	fmt.Println(len(m))
}
