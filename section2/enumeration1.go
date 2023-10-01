package section2

import "fmt"

func Enumeration1() {

	/*
		// 일정한 규칙에 따라 사용되는 상수라면 열거형으로 선언하면 편하게 사용 가능
		const (
			JAN = 1
			FEB = 2
			MAR = 3
			APR = 4
			MAY = 5
			JUN = 6
		)
	*/

	const (
		JAN = iota + 1
		FEB
		MAR
		APR
		MAY
		JUN
	)

	fmt.Println("JAN : ", JAN)
	fmt.Println("FEB : ", FEB)
	fmt.Println("MAR : ", MAR)
	fmt.Println("APR : ", APR)
	fmt.Println("MAY : ", MAY)
	fmt.Println("JUN : ", JUN)

	const (
		A = (iota + 1) * 10
		B
		_
		D
	)

	fmt.Println("A : ", A)
	fmt.Println("B : ", B)
	fmt.Println("D : ", D)

}
