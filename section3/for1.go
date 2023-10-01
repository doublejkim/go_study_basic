package section3

import "fmt"

func For1() {

	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
	fmt.Println("============================")

	/*
		무한 루프

		for {
			// ...
		}

	*/

	// 배열을 받아 range 로 순회
	loc := []string{"Seoul", "Incheon", "Busan"}
	for index, name := range loc {
		fmt.Println("index : ", index, ", name : ", name)
	}
	fmt.Println("============================")

	for _, name := range loc {
		fmt.Println("name : ", name)
	}
}
