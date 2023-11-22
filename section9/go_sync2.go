package section9

import (
	"fmt"
	"runtime"
	"sync"
)

type count2 struct {
	num   int
	mutex sync.Mutex
}

func (c *count2) increment2() {
	c.mutex.Lock() // 스레드들로부터 연산 보호
	c.num += 1
	c.mutex.Unlock()
}

func (c *count2) result2() {
	fmt.Println(c.num)
}

func GoSync2() {

	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count2{num: 0}
	done := make(chan bool)

	for i := 1; i < 10_000; i++ {
		go func() {
			c.increment2()
			done <- true
			runtime.Gosched() // cpu 양보
		}()
	}

	for i := 1; i < 10_000; i++ {
		<-done
	}

	c.result2()

}
