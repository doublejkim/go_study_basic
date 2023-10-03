package section5

import "fmt"

func Numeric2() {

	var num1 float32 = 0.14
	var num2 float32 = .765
	var num3 float32 = 123.0345123
	var num4 float32 = 10.0

	// 지수 표기법
	var num5 float32 = 14e6
	var num6 float64 = .156875e+3
	var num7 float64 = 5.32521e-10

	fmt.Println("num1 : ", num1)
	fmt.Println("num2 : ", num2)
	fmt.Println("num3 : ", num3)
	fmt.Println("num4 : ", num4)
	fmt.Println("num4 - 0.1 : ", num4-0.1)
	fmt.Println("float32(num4-0.1) : ", float32(num4-0.1))
	fmt.Println("float64(num4-0.1) : ", float64(num4-0.1))

	fmt.Println("num5 : ", num5)
	fmt.Println("num6 : ", num6)
	fmt.Println("num7 : ", num7)

}
