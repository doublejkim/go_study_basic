package section9

import "fmt"

func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 0; i < cnt; i++ {
			sum += i
		}
		tot <- sum
	}()
	return tot
}

func GoChannelEx2() {

	c := sum(100)
	fmt.Println("c : ", <-c)
}
