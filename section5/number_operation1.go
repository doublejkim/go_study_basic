package section5

import (
	"fmt"
	"math"
)

func Number_Operation1() {

	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("n1 : ", n1)
	fmt.Println("n2 : ", n2)
	fmt.Println("n3 : ", n3)
	fmt.Println("n4 : ", n4)

	n5 := 100_000 // int
	n6 := int16(10_000)
	n7 := uint8(100)

	//fmt.Println(n5 + n6) 에러 발생
	fmt.Println("n5 + int(n6)", n5+int(n6))
	fmt.Println("n6 + int16(n7)", n6+int16(n7))

	fmt.Println("=================================")

	var num1 uint16 = 123
	var num2 uint16 = 80

	fmt.Println("num1 + num2", num1+num2)
	fmt.Println("num1 - num2", num1-num2)
	fmt.Println("num1 * num2", num1*num2)
	fmt.Println("num1 / num2", num1/num2)
	fmt.Println("num1 % num2", num1%num2)
	fmt.Println("num1 << 2", num1<<2)
	fmt.Println("num1 >> 2", num1>>2)
	fmt.Println("^num1", ^num1)

	fmt.Println("=================================")

	var x1 uint8 = math.MaxUint8
	var x2 uint8 = x1 + 2
	fmt.Println(x2)
}
