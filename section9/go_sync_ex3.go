package section9

import (
	"fmt"
	"sync"
)

func GoSyncEx3() {

	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt += 1
			wg.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt -= 1
			wg.Done()
		}(i)
	}

	// Add == Done 횟수가 같아야함
	wg.Wait()
	fmt.Println("WaitGroup End! Cnt >>>>> : ", cnt)
}
