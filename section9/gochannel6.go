package section9

import "fmt"

func GoChannel6() {

	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- "Good~!!!"
		}
	}()

	val1, ok1 := <-ch
	fmt.Println("val1 : ", val1, ", ok1 : ", ok1)

	val2, ok2 := <-ch
	fmt.Println("val2 : ", val2, ", ok2 : ", ok2)

	val3, ok3 := <-ch
	fmt.Println("val3 : ", val3, ", ok3 : ", ok3)

}
