package section3

import "fmt"

func Switch3() {

	a := 30 / 15
	switch a {
	case 2, 4, 6, 8:
		fmt.Println("2, 4, 6, 8")
	case 1, 3, 5, 7, 9:
		fmt.Println("1, 3, 5, 7, 9")
	default:
		fmt.Println("Default!!!!")
	}

	switch e := "go"; e {
	case "java":
		fmt.Println("Java!!!")
		fallthrough
	case "go":
		fmt.Println("Go!!!")
		fallthrough
	case "python":
		fmt.Println("Python!!!")
		fallthrough
	case "ruby":
		fmt.Println("Ruby!!!")
		fallthrough
	case "c":
		fmt.Println("C!!!")
	}
}
