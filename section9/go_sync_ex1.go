package section9

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func onceTest() {
	// 한 번만 실행할 코드
	fmt.Println("Once Test!!!!!!!!!!!!!!")
}

func GoSyncEx1() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once)

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println("Goroutine : i : ", n)
			once.Do(onceTest)
		}(i)
	}

	time.Sleep(3 * time.Second)
}
