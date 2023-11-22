package section9

import (
	"fmt"
	"runtime"
)

type count1 struct {
	num int
}

func (c *count1) increment() {
	c.num += 1
}

func (c *count1) result() {
	fmt.Println(c.num)
}

func GoSync1() {

	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count1{num: 0}
	done := make(chan bool)

	for i := 1; i < 10_000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // cpu 양보
		}()
	}

	for i := 1; i < 10_000; i++ {
		<-done
	}

	c.result() // 실행할 마다 결과가 달라짐

}
