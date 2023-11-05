package section8

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Caculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func Struct1() {

	fmt.Println("=================== 구조체 선언과 사용")

	kim := Account{number: "111-111", balance: 10_000_000, interest: 0.015}
	lee := Account{number: "111-222", balance: 12_000_000}
	park := Account{number: "333-333", interest: 0.025}
	cho := Account{"555-555", 15_000_000, 0.03}

	fmt.Println(kim)
	fmt.Println(lee)
	fmt.Println(park)
	fmt.Println(cho)

	fmt.Println("=================== 구조체와 메소드 사용 ")

	fmt.Println("kim Calculate : ", int(kim.Caculate()))
	fmt.Println("lee Calculate : ", int(lee.Caculate()))
	fmt.Println("park Calculate : ", int(park.Caculate()))
	fmt.Println("cho Calculate : ", int(cho.Caculate()))
}
