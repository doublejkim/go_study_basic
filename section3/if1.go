package section3

import "fmt"

func If1() {

	var a int = 15
	b := 15

	if a >= 15 {
		fmt.Println("a 는 15 이상")
	}

	if b >= 30 {
		fmt.Println("b 는 30 이상")
	}

	if a >= 0 && a < 10 {
		fmt.Println("a >= 0 && a <10")
	} else if a >= 10 && a < 20 {
		fmt.Println("a >= 10 && a < 20 ")
	} else {
		fmt.Println("Unknown....")
	}

}
