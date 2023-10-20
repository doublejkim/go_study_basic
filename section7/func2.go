package section7

import "fmt"

func Func2() {
	fmt.Println("======================= Callback Test")

	CallBackTest1(11, 22, TestSum)

	fmt.Println("======================= 슬라이스에 함수 저장 및 활용")

	sliceFunc := []func(int, int) int{TestSum, TestMinus, TestMultiply, TestDivide}

	for idx, f := range sliceFunc {
		fmt.Println("idx : ", idx, " : ", f(9, 3))
	}

	fmt.Println("======================= 맵에 함수 저장 및 활용")

	mapFunc := map[string]func(int, int) int{
		"sum":      TestSum,
		"multiply": TestMultiply,
	}

	fmt.Println("call sum : ", mapFunc["sum"](10, 3))
	fmt.Println("call multiply : ", mapFunc["multiply"](10, 3))

	fmt.Println("======================= 익명 함수 사용")

	annonymousFunc := func(a int) int {
		return a * 2
	}

	fmt.Println("annonymousFunc : ", annonymousFunc(30))

}

func TestSum(a int, b int) int {
	return a + b
}

func TestMinus(a int, b int) int {
	return a - b
}

func TestMultiply(a int, b int) int {
	return a * b
}

func TestDivide(a int, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func CallBackTest1(num1 int, num2 int, f func(int, int) int) {
	fmt.Println("result call f : ", f(num1, num2))
}
