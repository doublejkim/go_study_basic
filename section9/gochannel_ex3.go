package section9

import "fmt"

func receiveOnly2(cnt int) <-chan int {
	sum := 1
	tot := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 777
		close(tot)
	}()
	return tot
}

func total(c <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a = a + i
		}
		tot <- a
	}()
	return tot
}

func GoChannelEx3() {

	c := receiveOnly2(100) // 채널 반환
	output := total(c)
	//output<- 777 // error. total 함수의 리턴타입은 수신전용 채널임
	fmt.Println("<-output : ", <-output)
}
