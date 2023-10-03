package section5

import "fmt"

func String2() {

	var str1 string = "Golang"
	var str2 string = "World"
	var str3 string = "고프로그래밍"

	// 각 문자의 코드값
	fmt.Println("str1[] : ", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Println("str2[] : ", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Println("str3[] : ", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	fmt.Printf("str1[] %c %c %c %c %c %c\n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Printf("str2[] : %c %c %c %c %c\n", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Printf("str3[] : %c %c %c %c %c %c\n", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5]) // ???

	fmt.Println("===================================")

	conStr := []rune(str3)
	fmt.Printf("conStr : %c %c %c %c %c %c\n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5])

	fmt.Println("===================================")

	for i, char := range str1 {
		fmt.Printf("str1 : %c(%d)\t", char, i)
	}

	fmt.Println("\n===================================")

	for i, char := range str2 {
		fmt.Printf("str2 : %c(%d)\t", char, i)
	}

	fmt.Println("\n===================================")

	for i, char := range str3 {
		fmt.Printf("str3 : %c(%d)\t", char, i)
	}
}
