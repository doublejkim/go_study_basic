package section10

import (
	"fmt"
	"log"
)

func notZero(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprintf("Hello Golang : %d", n)
		return s, nil
	}

	return "", fmt.Errorf("%d를 입력했음!!. 에러 발생 : ", n) //errors.New() 로 대체 가능
}

func Error2() {

	a, err := notZero(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result a : ", a)

	b, err := notZero(0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result b : ", b)
}
