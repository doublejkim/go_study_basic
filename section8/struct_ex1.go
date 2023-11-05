package section8

import "fmt"

type Account1 struct {
	number   string
	balance  float64
	interest float64
}

// 생성자 패턴
func NewAccount(number string, balance float64, interest float64) *Account {
	return &Account{number, balance, interest}
}

func StructEx1() {

	account1 := NewAccount("111-222", 1_200_000, 0.04)
	fmt.Printf("account1 : %#v\n", account1)
}
