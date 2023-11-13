package section9

import (
	"fmt"
	"time"
)

func GoChannel3() {
	//

	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}
}
