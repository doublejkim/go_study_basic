package section6

import "fmt"

func SliceEx1() {

	fmt.Println("======================= 슬라이스 - append")

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{11, 12, 13, 14, 15}
	s3 := []int{21, 22, 23, 24, 25}

	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)
	s3 = append(s2, s3[0:3]...)

	fmt.Println("s1 : ", s1)
	fmt.Println("s2 : ", s2)
	fmt.Println("s3 : ", s3)

	fmt.Println("======================= 슬라이스 -- loop 안에서 append")

	s4 := make([]int, 0, 5)
	//s4 := []int{}  // capa 를 지정하지 않을경우 최초 append 되는시점에 capa 가 설정됨. 아래의 예는 1개..
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("s4 - len : %d, cap : %d, value : %v\n", len(s4), cap(s4), s4)
	}
}
