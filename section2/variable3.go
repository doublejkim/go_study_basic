package section2

import "fmt"

func Variable3() {

	shortVar1 := 5
	shortVar2 := "Test"
	shortVar3 := true

	fmt.Println("shortVar1 : ", shortVar1)
	fmt.Println("shortVar2 : ", shortVar2)
	fmt.Println("shortVar3 : ", shortVar3)

	if i := 10; i < 11 {
		fmt.Println("Short Variable in if test")
	}

}
