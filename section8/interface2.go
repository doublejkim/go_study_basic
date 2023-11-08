package section8

import "fmt"

type Dog struct {
	name   string
	weight int
}

// bite method 구현
func (d Dog) bite1() {
	fmt.Println(d.name, " bites!!")
}

type Behavior interface {
	bite1()
}

func Interface2() {

	fmt.Println("=================== 인터페이스 사용 1")

	dog1 := Dog{"poll", 10}
	var interface1 Behavior
	interface1 = dog1
	interface1.bite1()

	fmt.Println("=================== 인터페이스 사용 2")

	dog2 := Dog{"marry", 13}
	interface2 := Behavior(dog2)
	interface2.bite1()

	fmt.Println("=================== 인터페이스 사용 - 배열로 사용해보기 1")
	inters := []Behavior{dog1, dog2}

	for idx, _ := range inters {
		inters[idx].bite1()
	}

	fmt.Println("=================== 인터페이스 사용 - 배열로 사용해보기 2")
	for _, val := range inters {
		val.bite1()
	}
}
