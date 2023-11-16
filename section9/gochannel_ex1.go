package section9

import (
	"fmt"
	"time"
)

func sendOnly(c chan<- int, cnt int) {

	for i := 0; i < cnt; i++ {
		c <- i
	}

	c <- 777

	//fmt.Println(<-c) // 에러 발생

}

func receiveOnly(c <-chan int) {
	for i := range c {
		fmt.Println("received : ", i)
	}

	fmt.Println(<-c)
}

func GoChannelEx1() {

	c := make(chan int)

	go sendOnly(c, 10) // 발신 전용
	go receiveOnly(c)  // 수신 전용

	time.Sleep(2 * time.Second)
}
