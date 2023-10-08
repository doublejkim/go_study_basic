package section6

import "fmt"

func Array2() {

	fmt.Println("======================= 배열 순회")

	arr1 := [5]int{1, 10, 100, 1000, 10000}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(i, " : ", arr1[i])
	}

	fmt.Println("======================= 배열 순회 range 활용")

	arr2 := [5]int{10, 20, 30, 40, 50}

	for i, v := range arr2 {
		fmt.Println(i, " : ", v)
	}

	fmt.Println("======================= 배열 순회 range 활용, 값 만")

	for _, v := range arr2 {
		fmt.Println(v)
	}
}
