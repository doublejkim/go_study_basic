package section9

import (
	"fmt"
	"sync"
)

func GoSyncEx2() {

	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("WaitGroup : ", n)
			wg.Done()
		}(i)
	}

	// Add == Done 횟수가 같아야함
	wg.Wait()

}
