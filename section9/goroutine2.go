package section9

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func exe(name int) {
	r := rand.Intn(100)
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, " >>>>> ", r, i)
	}
	fmt.Println(name, " end : ", time.Now())
}

func Goroutine2() {

	runtime.GOMAXPROCS(runtime.NumCPU())                        // 현재 시스템의 cpu 코어 개수 반환후 설정
	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(0)) // 설정값 출력

	fmt.Println("Main Routine Start : ", time.Now())
	for i := 0; i < 100; i++ {
		go exe(i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Main Routine End : ", time.Now())
}
