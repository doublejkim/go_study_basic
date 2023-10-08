package section6

import "fmt"

func Array3() {

	fmt.Println("======================= 값 복사 확인")

	arr1 := [5]int{10, 20, 30, 40, 50}
	arr2 := arr1

	fmt.Println("Println() arr1 : ", arr1, &arr1)
	fmt.Println("Println() arr2 : ", arr2, &arr2)

	fmt.Printf("arr1 : Addr(%p), Value(%v)\n", &arr1, arr1)
	fmt.Printf("arr2 : Addr(%p), Value(%v)\n", &arr2, arr2)

}
