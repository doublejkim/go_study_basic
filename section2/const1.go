package section2

import "fmt"

func Const1() {

	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10

	//a = "xxx" // 다른 값을 대입하려고 하면 에러 발생
	//c = GetHeight() // 함수 등을 통해 초기화 불가 (해당 함수 내에서 mutable 한 값이 발생할 수 있기 때문..)

	fmt.Println("a : ", a, ", b : ", b, ", c : ", c)
}
