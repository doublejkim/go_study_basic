package section5

import (
	"fmt"
	"strings"
)

func StringOper3() {

	// 일반연산
	str1 := "aaaaa bbbbb ccccc" +
		" xxxxxxx" +
		" zzzzzzzz"

	str2 := "qqqqwwwwwqqqqqqqqqqqq"

	fmt.Println("ex 1 : ", str1+str2)

	// Join

	strSet := []string{} // 슬라이스 선언
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println("ex 2 : ", strings.Join(strSet, ""))

}
