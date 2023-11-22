package section9

import (
	"fmt"
	"time"
)

func GoSync3() {

	data := 0

	go func() {
		for i := 1; i <= 10; i++ {
			data += 1
			fmt.Println("Write : data : ", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read 1 : data : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read 2 : data : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
}
