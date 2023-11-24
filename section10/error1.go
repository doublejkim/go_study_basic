package section10

import (
	"fmt"
	"log"
	"os"
)

func Error1() {

	f, err := os.Open("unnamedFile")

	if err != nil {
		//fmt.Println("error Msg : ", err.Error())
		log.Fatal("error Msg2 : ", err.Error())
	}

	fmt.Println(f.Name())

}
