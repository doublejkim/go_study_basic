package section9

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func GoSync4() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.Lock()
			data += 1
			fmt.Println("Write : data : ", data)
			time.Sleep(200 * time.Millisecond)
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.RLock()
			fmt.Println("Read 1 : data : ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.RLock()
			fmt.Println("Read 2 : data : ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()

	time.Sleep(15 * time.Second)
}
