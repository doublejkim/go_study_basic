package section10

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error : ", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Open File Name : ", f.Name())
	}

	defer f.Close()
}

func PanicRecover4() {

	fileOpen("test")

	fmt.Println("End Main!!!!")
}
