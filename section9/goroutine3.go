package section9

import (
	"fmt"
	"runtime"
	"time"
)

func Goroutine3() {

	runtime.GOMAXPROCS(1) // 1개만 사용하겠다고 선언

	s := "Goroutine Closure"

	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) // 반복문 클로저는 일반적으로 즉시 실행. But, 고루틴 클로저는 가장 나중에 실행 (반복문종료 후)
	}

	time.Sleep(3 * time.Second)
}
