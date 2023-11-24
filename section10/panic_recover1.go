package section10

import "fmt"

func PanicRecover1() {
	fmt.Println("Start !!!")
	panic("Error occurred : user stopped!!")
	fmt.Println("Finish !!!") // 실행 불가
}
