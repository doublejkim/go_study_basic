package section6

import "fmt"

func Pointer1() {

	fmt.Println("======================= 포인터 변수 선언시 초기화 할때와 명시적 new 할때의 비교")

	var a *int            // nil
	var b *int = new(int) // 임의의 메모리 공간이 할당 됨

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)

	fmt.Println("======================= 포인터 변수로 다른 변수의 주소를 참조 예제")

	i := 7

	p1 := &i
	p2 := &i

	fmt.Println("i : ", i, ", &i : ", &i)
	fmt.Println("*p1 : ", *p1, ", p1 : ", p1)
	fmt.Println("*p1 : ", *p2, ", p1 : ", p2)

	*p2 = 123 // p2 포인터 변수가 레퍼런스하고있는 메모리주소에 저장된 값을 변경해보기

	fmt.Println("i : ", i, ", &i : ", &i)
	fmt.Println("*p1 : ", *p1, ", p1 : ", p1)
	fmt.Println("*p1 : ", *p2, ", p1 : ", p2)
}
