package section8

import "fmt"

type totCost func(int, int) int

func describe(cnt int, price int, fn totCost) {
	fmt.Printf("cnt : %d, price : %d, totalCost : %d", cnt, price, fn(cnt, price))
}

func UserStruct2() {
	var orderPrice totCost
	orderPrice = func(cnt int, price int) int {
		return (cnt * price) + 100_000
	}

	describe(3, 5000, orderPrice)
}
