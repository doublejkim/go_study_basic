package section3

import "fmt"

func For2() {

	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println("sum1 : ", sum1)

	fmt.Println("============================")

	sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
		i++ // 참고 : golang 에서의 후치연산은 반환값이 없어서 후치연산 이후 대입하면 에러발생. ex : j := i++
	}
	fmt.Println("sum2 : ", sum2)

	fmt.Println("============================")

	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("sum3 : ", sum3)

	fmt.Println("============================")

	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("i : ", i, ", j : ", j)
	}

}
