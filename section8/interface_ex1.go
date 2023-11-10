package section8

import "fmt"

func printContents(v interface{}) {
	fmt.Printf("Type : (%T)\n", v)
	fmt.Println("v : ", v)
}

func InterfaceEx1() {

	var a interface{}

	a = 7.5
	printContents(a)

	a = "Golang.."
	printContents(a)

	a = true
	printContents(a)

	a = nil
	printContents(a)

}
