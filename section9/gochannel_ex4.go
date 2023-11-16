package section9

import (
	"fmt"
	"time"
)

func GoChannelEx4() {

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 11
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Hello"
			time.Sleep(700 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("ch1 : ", num)
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
			}
		}
	}()

	time.Sleep(10 * time.Second)

}
