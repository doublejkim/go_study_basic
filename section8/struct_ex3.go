package section8

import "fmt"

type Employee struct {
	name   string
	salary int
	bonus  int
}

func (e Employee) Calculate() int {
	fmt.Println("Employee's Calculate called!!")
	return e.salary + e.bonus
}

func (e Executives) Calculate() int {
	fmt.Println("Executives's Calculate called!!")
	return e.salary + e.bonus + e.specialBonus
}

type Executives struct {
	Employee
	specialBonus int
}

func StructEx3() {
	// 구조체 임베디드 패턴
	// 다른 관점으로 메소드를 재 사용하는 장점 제공
	// 상속을 허용하지 않는 Go 언어에서 메소드 상속 활용을 위한 패턴

	ep1 := Employee{"kim", 2_000_000, 150_000}
	ep2 := Employee{"park", 1_500_000, 200_000}

	ex1 := Executives{
		Employee{"lee", 5_000_000, 150_000},
		200_000,
	}

	fmt.Println("ep1 : ", ep1.Calculate())
	fmt.Println("ep2 : ", ep2.Calculate())

	fmt.Println("ex1 : ", ex1.Calculate()+ex1.specialBonus)
}

func StructEx4() {

	ep1 := Employee{"kim", 2_000_000, 150_000}
	ep2 := Employee{"park", 1_500_000, 200_000}

	ex1 := Executives{
		Employee{"lee", 5_000_000, 150_000},
		200_000,
	}

	fmt.Println("ep1 : ", ep1.Calculate())
	fmt.Println("ep2 : ", ep2.Calculate())

	fmt.Println("ex1 : ", ex1.Calculate()) // Executives 리시버를 사용하는 메소드가 호출됨
}
