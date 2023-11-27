package section11

import (
	"fmt"
	"os"
)

func FileWrite3() {

	file, err := os.Open("files/sample.txt")
	if err != nil {
		fmt.Println("File Open Error : ")
		return
	}

	defer file.Close()

	fileInfo, err := file.Stat() // 파일 사이즈 확인 위해 정보 획득
	if err != nil {
		fmt.Println("File Stat Error")
		return
	}

	fd1 := make([]byte, fileInfo.Size()) // 슬라이스에 읽은 내용 담음
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력 (1) : ", fileInfo)
	fmt.Println("파일 이름 (2) : ", fileInfo.Name())
	fmt.Println("파일 크기 (3) : ", fileInfo.Size())
	fmt.Println("파일 수정 시간 (4) : ", fileInfo.ModTime())
	fmt.Printf("읽기 작업 (1) 완료 (%d bytes)\n\n", ct1)
	//fmt.Println(fd1)
	fmt.Println(string(fd1)) // 타입 변환하여 코드값을 string 으로..

	fmt.Println("========================================")

	o1, err := file.Seek(10, 0) // whence : 0 : 처음위치, 1 : 현재 위치, 2 : 마지막 위치
	if err != nil {
		fmt.Println("File Seek Error")
		return
	}

	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	if err != nil {
		fmt.Println("File Read Error")
		return
	}

	fmt.Printf("읽기 작업 (2) 완료 (%d bytes) (%d ret)\n\n", ct2, o1)
	fmt.Println(string(fd2))
	fmt.Println("========================================")

	o2, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println("File Seek Error")
		return
	}
	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 3) // offset 위치부터 읽어옴
	if err != nil {
		fmt.Println("File ReadAt Error")
		return
	}
	fmt.Printf("읽기 작업 (3) 완료 (%d bytes) (%d ret)\n\n", ct3, o2)
	fmt.Println(string(fd3))
	fmt.Println("========================================")

}
