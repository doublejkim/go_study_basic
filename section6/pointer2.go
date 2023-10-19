package section6

import "fmt"

func Pointer2() {

	i := 7
	p := &i

	fmt.Println("i : ", i, ", *p : ", *p, ", &i : ", &i, ", p : ", p)

	*p++ // 포인터 변수 p 가 레퍼런스하고있는 변수의 메모리주소에 저장되어있는 실제값(역참조) 을 증가시킴

	fmt.Println("i : ", i, ", *p : ", *p, ", &i : ", &i, ", p : ", p)
}
