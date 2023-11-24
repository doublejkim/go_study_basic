package section10

import "fmt"

func runFunc2() {
	defer func() {
		s := recover()
		fmt.Println("Error Message : ", s)
	}()

	panic("Error occurred!!!")
}

func PanicRecover2() {

	runFunc2()

	fmt.Println("Hello???")
}
