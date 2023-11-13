package section9

import (
	"fmt"
	"runtime"
)

func GoChannel4() {
	// 비동기 : 버퍼사용

	runtime.GOMAXPROCS(1)
	ch := make(chan bool, 2) // 채널 버퍼의 용량을 설정 가능.
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}
}
