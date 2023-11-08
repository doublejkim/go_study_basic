package section8

import "fmt"

type testInterface1 interface{}

func Interface1() {
	var t testInterface1
	fmt.Println("testInterface1 : ", t)
}
