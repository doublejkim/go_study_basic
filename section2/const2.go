package section2

import "fmt"

func Const2() {

	const a, b int = 10, 99
	const c, d, e = true, 0.12, "test"
	const (
		x, y     int16 = 10, 25
		data, no       = "Data", 777
	)

	fmt.Println("a : ", a, ", b : ", b, ", c :", c)
	fmt.Println("x : ", x, ", y : ", y, ", data :", data, "no : ", no)

}
