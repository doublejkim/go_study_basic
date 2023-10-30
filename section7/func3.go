package section7

import "fmt"

func ex_f1() {
	fmt.Println("f1 : start!")
	defer ex_f2()
	fmt.Println("f1 : end!")
}

func ex_f2() {
	fmt.Println("f2 : called!!!!!!")
}

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("Hi!!!")
	}()
}

func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("ex - i : ", i)
	}
}

func start(t string) string {
	fmt.Println("start : ", t)
	return t
}

func end(t string) {
	fmt.Println("end : ", t)
}

func deferTest() {
	defer end(start("b")) // defer 를 사용하고 중첩함수를 사용하면 기대와 다르게 호출되는 경우를 놓칠 수 있음
	fmt.Println("in a")
}

func Func3() {

	fmt.Println("==================== ex_f1()")
	ex_f1()

	fmt.Println("==================== sayHello()")
	sayHello("Golang!!!!")

	fmt.Println("==================== stack()")
	stack()

	fmt.Println("==================== deferTest()")
	deferTest()
}
