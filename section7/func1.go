package section7

import "fmt"

func Func1() {

	fmt.Println("======================= 함수의 선언과 사용")

	result1, result2 := PlusMinusTest(5)

	fmt.Println("result1 : ", result1, ", result2 : ", result2)

	result3 := SumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("result3 : ", result3)

	slice1 := []int{5, 6, 7, 8}
	result4 := SumAll(slice1...)

	fmt.Println("result4 : ", result4)

}

func PlusMinusTest(num int) (result1 int, result2 int) {
	result1 = num + 1
	result2 = num - 1
	return
}

func SumAll(num ...int) int {
	total := 0
	for _, value := range num {
		total += value
	}
	return total
}
