package section5

import "fmt"

func Bool1() {

	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println("b1 : ", b1)
	fmt.Println("b2 : ", b2)
	fmt.Println("b3 : ", b3)

	fmt.Println("1 == 2", 1 == 2)
	fmt.Println("10 == 10", 10 == 10)
	fmt.Println("15 < 20", 15 < 20)
	fmt.Println("b2 && b3", b2 && b3)
}
