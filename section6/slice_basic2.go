package section6

import "fmt"

func SliceBasic2() {

	fmt.Println("======================== 배열의 값 복사 (슬라이스와의 비교용)")

	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1
	arr2[0] = 7

	fmt.Println("arr1 : ", arr1)
	fmt.Println("arr2 : ", arr2)

	fmt.Println("======================== 슬라이스 참조 복사 (슬라이스와의 비교용)")

	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1
	slice1[0] = 7

	fmt.Println("slice1 : ", slice1)
	fmt.Println("slice2 : ", slice2)

	fmt.Println("======================== 슬라이스 순회")

	for i, v := range slice1 {
		fmt.Println("i : ", i, "v : ", v)
	}

	fmt.Println("======================== 슬라이스 예외 상황 ")

	slice3 := make([]int, 50, 100) // capa 는 100 이지만 길이(50) 만큼 0~49 까지만 초기화 됨
	fmt.Println("slice3[4] : ", slice3[4])
	// fmt.Println("slice3[50] : ", slice3[50]) // panic 발생 : index out of range // 길이만큼만 초기화됨

}
