package section9

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func GoSync5() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) // 비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Waiting... : ", n)
			condition.Wait()
			fmt.Println("Waiting End : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
	}

	/*for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println("Wake Goroutine(Signal) : ", i)
		condition.Signal() // 한개씩 깨움 (모든 고루틴 생성 후)
		mutex.Unlock()
	}*/

	mutex.Lock()
	fmt.Println("Wake Goroutine(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)
}
