package section5

import (
	"fmt"
	"unicode/utf8"
)

func String1() {

	var str1 string = "c:\\temp\\downlaod"
	str2 := `c:\temp\download` // 백스쿼트를 사용하면 escape 를 사용하지 않아도 됨

	fmt.Println("str1 : ", str1)
	fmt.Println("str2 : ", str2)

	fmt.Println("===================================")

	var str3 string = "Hello, World"
	var str4 string = "안녕하세요"
	var str5 string = "\ud55c\uae00"

	fmt.Println("str3 : ", str3)
	fmt.Println("str4 : ", str4)
	fmt.Println("str5 : ", str5)

	fmt.Println("===================================")

	// 길이(바이트 수)
	fmt.Println("str3 len : ", len(str3))
	fmt.Println("str4 len : ", len(str4))

	// 길이(실제 길이)
	fmt.Println("str3 utf8.RuneCountInString() : ", utf8.RuneCountInString(str3))
	fmt.Println("str4 utf8.RuneCountInString() : ", utf8.RuneCountInString(str4))
	fmt.Println("str4 len([]rune(str4)", len([]rune(str4))) // 이렇게도 실제 길이를 구할 수 있지만. 많이 사용하지는 않음

}
