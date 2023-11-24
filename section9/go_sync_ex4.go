package section9

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func GoSyncEx4() {
	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			//cnt += 1
			atomic.AddInt64(&cnt, 1)
			wg.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			//cnt -= 1
			atomic.AddInt64(&cnt, -1)
			wg.Done()
		}(i)
	}

	// Add == Done 횟수가 같아야함
	wg.Wait()
	finalCnt := atomic.LoadInt64(&cnt) // cnt 를 가져와도되지만 보통 이런식으로 다시 로드함
	fmt.Println("WaitGroup End! Cnt >>>>> : ", cnt, ", finalCnt : ", finalCnt)
}
