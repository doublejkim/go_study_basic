package section8

import (
	"fmt"
	"reflect"
)

func InterfaceEx2() {

	fmt.Println("=========== Type Assertion 테스트")

	var a interface{} = 15

	b := a
	c := a.(int)
	//d := a.(float64) // panic : interface conversion

	fmt.Println("a : ", a)
	fmt.Println("reflect.TypeOf(a)", reflect.TypeOf(a))
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	//fmt.Println("d : ", d)

	if v, ok := a.(int); ok {
		fmt.Println("conversion 1 v : ", v)
	}

	if v, ok := a.(float64); ok {
		fmt.Println("conversion 2 v : ", v)
	}
}

func checkType(arg interface{}) {

	switch arg.(type) {
	case bool:
		fmt.Println("This is a bool : ", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int : ", arg)
	case float32, float64:
		fmt.Println("This is a float : ", arg)
	case string:
		fmt.Println("This is a string : ", arg)
	case nil:
		fmt.Println("This is a nil : ", arg)
	default:
		fmt.Println("Unknown type... : ", arg)
	}
}

func InterfaceEx3() {
	checkType(true)
	checkType(12)
	checkType(12.3)
	checkType("what?")
	checkType(nil)
}
