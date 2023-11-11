package section9

import (
	"fmt"
	"time"
)

func exe1(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Println("exe 1 [", i, "] : ", time.Now())
		time.Sleep(500 * time.Millisecond)
	}
}

func exe2(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Println("exe 2 [", i, "] : ", time.Now())
		time.Sleep(500 * time.Millisecond)
	}
}

func exe3(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Println("exe 3 [", i, "] : ", time.Now())
		time.Sleep(500 * time.Millisecond)
	}
}

func Goroutine1() {

	cnt := 10
	go exe1(cnt)
	go exe2(cnt)
	exe3(cnt)
}
