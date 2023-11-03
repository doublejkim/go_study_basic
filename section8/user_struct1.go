package section8

import (
	"fmt"
	"strconv"
)

type Car struct {
	name  string
	color string
	price int
	tax   int
}

// 구조체 <-> 메소드 바인딩
// 일반 메소드
func Price(c Car) int {
	return c.price + c.tax
}

func (c Car) Price() string {
	return c.name + "'s Price : " + strconv.Itoa(c.price+c.tax)
}

func (c *Car) SetName(newName string) {
	c.name = newName
}

func UserStruct1() {

	bmw := Car{name: "520d", price: 50000000, color: "white", tax: 5000000}
	benz := Car{name: "22d", price: 60000000, color: "black", tax: 6000000}

	fmt.Println("bmw : ", bmw, &bmw)
	fmt.Println("benz : ", benz, &benz)

	fmt.Println("bmw Price : ", Price(bmw))
	fmt.Println("benz Price : ", Price(benz))

	fmt.Println("ex - ", bmw.Price())
	fmt.Println("ex - ", benz.Price())

	bmw.SetName("아반떼")
	fmt.Println("ex - ", bmw.Price())

}
