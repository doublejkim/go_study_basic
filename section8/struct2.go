package section8

import "fmt"

func Struct2() {

	// 선언 방법1
	var kim *Account = new(Account)
	kim.number = "111-111"
	kim.balance = 10_000_000
	kim.interest = 0.015

	// 선언 방법2
	hong := &Account{number: "222-222", balance: 15_000_000, interest: 0.04}

	// 선언 방법3
	lee := new(Account)
	lee.number = "333-333"
	lee.balance = 130_000_000
	lee.interest = 0.025

	fmt.Println("kim : ", kim)
	fmt.Println("hong : ", hong)
	fmt.Println("lee : ", lee)

	fmt.Printf("kim 2 : %#v\n", kim)
	fmt.Printf("hong 2 : %#v\n", hong)
	fmt.Printf("lee 2 : %#v\n", lee)

	fmt.Println("kim Calculate : ", int(kim.Caculate()))
	fmt.Println("hong Calculate : ", int(hong.Caculate()))
	fmt.Println("lee Calculate : ", int(lee.Caculate()))
}
