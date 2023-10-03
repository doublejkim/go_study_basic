package section5

import "fmt"

func Numeric1() {

	var num1 int = 17
	var num2 int = -68
	var num3 int = 0631
	var num4 int = 0x32fa2c75

	fmt.Println("num1 : ", num1)
	fmt.Println("num2 : ", num2)
	fmt.Println("num3 : ", num3)
	fmt.Println("num4 : ", num4)

	fmt.Println("=========================")

	var char1 byte = 72 // byte == unit8
	var char2 byte = 0110
	var char3 byte = 0x48

	fmt.Printf("%c %c %c\n", char1, char2, char3)
	fmt.Printf("%d %d %d\n", char1, char2, char3)
	fmt.Printf("%d %o %x\n", char1, char2, char3)

	fmt.Println("=========================")

	var char4 rune = 50556   // 유니코드 10진수
	var char5 rune = 0142574 // 44032 (8진수)
	var char6 rune = 0xC57C  // 44032(16진수)

	fmt.Printf("%c %c %c\n", char4, char5, char6)
	fmt.Printf("%d %d %d\n", char4, char5, char6)
	fmt.Printf("%d %o %x\n", char4, char5, char6)

}
