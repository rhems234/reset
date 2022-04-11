package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan string) {
	var msg string
	fmt.Scanln(&msg)
	ch <- msg
}

func main() {
	ch := make(chan string)

	defer func() {
		close(ch)
	}()

	go sendMessage(ch)

	for t := 10; 0 < t; t-- {
		select {
		case ch1 := <-ch:
			fmt.Printf("%s 메세지가 발송되었습니다.", ch1)
		default:
			fmt.Printf("%d초 안에 메세지를 입력하세요.\n", t)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("메세지 발송에 실패했습니다.")
}
