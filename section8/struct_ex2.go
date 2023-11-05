package section8

import "fmt"

type Account2 struct {
	number   string
	balance  float64
	interest float64
}

func CaculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CaculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func StructEx2() {

	kim := Account{"111-111", 1_000_000, 0.015}
	lee := Account{"222-222", 2_000_000, 0.015}

	fmt.Printf("kim : %#v\n", kim)
	fmt.Printf("lee : %#v\n", lee)

	CaculateD(kim)
	CaculateP(&lee)

	fmt.Printf("kim : %#v\n", kim)
	fmt.Printf("lee : %#v\n", lee)
}
