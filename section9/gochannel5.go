package section9

import "fmt"

func GoChannel5() {

	// Close : 채널 닫기. 주의 : 닫힌 채널에 값 전송시 panic 발생
	// Range : 채널 안에서 값을 꺼냄 (순회), 채널 닫아야(Close) 반복문 정료 -> 채널이 열려있고 값 전송하지 않으면 계속 대기

	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch) // 5회 채널에 값 전송 후 채널 닫기
	}()

	for i := range ch { // 채널에서 값을 꺼내옴 (채널 close 될때 까지)
		fmt.Println("i : ", i)
	}
}
