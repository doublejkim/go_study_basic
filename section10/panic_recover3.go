package section10

import "fmt"

func runFunc3() {
	defer func() {
		s := recover()
		fmt.Println("Error Message : ", s)
	}()

	a := [3]int{1, 2, 3}

	for i := 0; i < 5; i++ {
		fmt.Println("ex : ", a[i]) // 에러 발생
	}

}

func PanicRecover3() {

	runFunc3()

	fmt.Println("Hello???")
}
