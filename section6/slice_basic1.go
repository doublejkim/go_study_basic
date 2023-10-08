package section6

import "fmt"

func SliceBasic1() {

	fmt.Println("======================== 슬라이스 선언과 사용")

	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	slice3[4] = 10 // 값 수정 가능

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	fmt.Println("======================== 슬라이스 선언과 사용 - make")

	var slice5 []int = make([]int, 5, 10) // length 는 5, capacity 는 10으로 확보 .
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)

	slice6[2] = 7

	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

	fmt.Println("======================== 슬라이스 선언과 사용 - nil")
	var slice9 []int // nil 슬라이스 -> make 로 할당한 후에 사용 가능

	if slice9 == nil {
		fmt.Println("This is Nil Slice. : ", slice9)
	}
}
