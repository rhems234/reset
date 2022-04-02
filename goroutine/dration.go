package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r := rand.Intn(3) // 0부터 3까지 난수 생성
	time.Sleep(time.Duration(r) * time.Second)
	// 난수를 Dration형으로 형변환 후 second로 계산
	fmt.Println(n)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i) // 동기 실행(주어진 일을 순서대로)
	}
	fmt.Scanln()
}
