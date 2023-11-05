package section8

import "fmt"

func Struct3() {

	fmt.Println("=================== 구조체 익명 선언")

	// 구조체 익명 선언
	car1 := struct{ name, color string }{"520d", "red"}

	fmt.Printf("car1 : %#v\n", car1)

	fmt.Println("=================== 구조체 익명선언 배열 활용")

	cars := []struct{ name, color string }{{"520d", "white"}, {"k3", "black"}, {"carnival", "gray"}}

	for _, c := range cars {
		fmt.Println("name : ", c.name, ", color : ", c.color)
	}

}
