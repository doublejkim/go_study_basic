package section6

import "fmt"

func SliceEx3() {

	fmt.Println("======================= 슬라이스 copy")

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println("slice2 : ", slice2)
	fmt.Println("slice3 : ", slice3) // slice3 는 초기화가 되지 않은 상태라서 복사되지 않음

	fmt.Println("======================= 슬라이스 copy")

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)

	b[0] = 7
	b[4] = 10

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)

	fmt.Println("======================= 슬라이스 - 슬라이싱(substr)할때 주의사향")

	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] // substring 으로 가져오면 참조타입으로 가져옴

	d[1] = 7

	fmt.Println("c : ", c)
	fmt.Println("d : ", d)

	fmt.Println("======================= 슬라이스")

	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7]

	fmt.Println("e : len : ", len(e), ", cap : ", cap(e), ", value : ", e)
	fmt.Println("f : len : ", len(f), ", cap : ", cap(f), ", value : ", f)

}
