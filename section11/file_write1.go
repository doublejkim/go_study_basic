package section11

import (
	"fmt"
	"os"
)

func FileWrite1() {

	file, err := os.Create("test_write1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	str := "Hello World..."

	n, err := file.Write([]byte(str)) // str 을 []byte 슬라이스로 변환. str 을 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, " 바이트 저장 완료")

}
