package section7

import "fmt"

func Closure1() {

	fmt.Println("===================== Closure 기본적인 사용")

	multiply := func(x, y int) int {
		return x * y
	}

	r1 := multiply(5, 10)

	fmt.Println("r1 : ", r1)

	m, n := 5, 10
	sum := func(c int) int {
		return m + n + c
	}

	r2 := sum(10)
	fmt.Println("r2 : ", r2)

	fmt.Println("===================== Closure 기본적인 사용")

	cntFunc := increaseCnt()

	fmt.Println("test ", cntFunc())
	fmt.Println("test ", cntFunc())
	fmt.Println("test ", cntFunc())
	fmt.Println("test ", cntFunc())
	fmt.Println("test ", cntFunc())
}

func increaseCnt() func() int {
	n := 0 // 지역변수 (캡쳐됨)
	return func() int {
		n += 1
		return n
	}
}
