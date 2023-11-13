package section9

import (
	"fmt"
	"time"
)

func work1(v chan int) {
	fmt.Println("Work1 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : E ---> ", time.Now())
	v <- 1
}

func work2(v chan int) {
	fmt.Println("Work2 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : E ---> ", time.Now())
	v <- 2
}

func GoChannel1() {

	fmt.Println("myFunc : S ---> ", time.Now())

	//var c chan int
	//c = make(chan int) // 이렇게 사용해도 문제 없음
	v := make(chan int) // int 형 채널 선언

	go work1(v)
	go work2(v)

	<-v
	<-v
	fmt.Println("myFunc : E ---> ", time.Now())

}
