package section3

import "fmt"

func Switch1() {

	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	switch b := 27; {
	case a < 0:
		fmt.Println(b, "는 음수")
	case a == 0:
		fmt.Println(b, "는 0")
	case a > 0:
		fmt.Println(b, "는 양수")
	}

	switch c := "go"; c {
	case "go":
		fmt.Println("Go!!!")
	case "java":
		fmt.Println("Java!!!")
	default:
		fmt.Println("Default!!!!!")
	}

	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("GoLang!!!")
	case "java":
		fmt.Println("Java!!!")
	default:
		fmt.Println("Default!!!!!")
	}

	//i := 20
	//j := 30
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println(i, " 는 ", j, " 보다 작다")
	case i == j:
		fmt.Println(i, " 와 ", j, " 는 같다")
	case i > j:
		fmt.Println(i, " 가 ", j, " 보다 크다")
	default:
		fmt.Println("Default!!!")
	}
}
