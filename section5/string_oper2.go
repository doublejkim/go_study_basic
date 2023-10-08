package section5

import "fmt"

func StringOper2() {

	str1 := "Golang"
	str2 := "World"

	fmt.Println("str1 == str2 : ", str1 == str2)
	fmt.Println("str1 != str2 : ", str1 != str2)
	fmt.Println("str1 < str2 : ", str1 < str2) // 아스키 번호 기준으로 앞에서부터 비교로 true
	fmt.Println("str1 > str2 : ", str1 > str2)

}
