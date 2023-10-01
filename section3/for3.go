package section3

import "fmt"

func For3() {

Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				fmt.Println("Now!! i : ", i, ", j : ", j)
				break Loop1 // 가장 바깥 쪽 for 문을 벗어남
			}
		}
	}

	fmt.Println("==================================")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // 짝 수 일때는 for 문 안의 다른 문장은 실행하지 않고 증감식 으로   돌아감
		}
		fmt.Println("i : ", i) //
	}

	fmt.Println("==================================")

Loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("i : ", i, ", j : ", j)
		}
	}

}
