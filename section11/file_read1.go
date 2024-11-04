package section11

import (
	"fmt"
	"os"
)

func FileRead1() {

	fmt.Println("============================= Open/Read 예제")

	file, err := os.Open("files/sample.txt")
	if err != nil {
		fmt.Println(" 1) Open Error")
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("1) Stat Error")
		return
	}

	fd1 := make([]byte, fileInfo.Size()) // 파일 크기 만큼 바이트 슬라이스 생성
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력 : ", fileInfo, "\n")
	fmt.Println("파일 이름 : ", fileInfo.Name(), "\n")
	fmt.Println("파일 크기 : ", fileInfo.Size(), "\n")
	fmt.Println("파일 수정 시간 : ", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업 완료 (1) (%d bytes)\n\n", ct1)
	//fmt.Println(fd1) // byte 값 그대로 출력됨
	fmt.Println(string(fd1))

	fmt.Println("============================= Seek 기본 예제")

	o1, err := file.Seek(20, 0) // whence - 0: 처음 위치, 1: 현재 위치, 2: 마지막 위치
	if err != nil {
		fmt.Println("2) Seek Error")
		return
	}

	fd2 := make([]byte, 20) // 앞에서 offset 20 만큼 이동한상태에서 20만큼 read
	ct2, err := file.Read(fd2)
	if err != nil {
		fmt.Println("2) Read Error")
		return
	}

	fmt.Printf("읽기 작업 완료 (2) (%d bytes) (%d ret)\n\n", ct2, o1)
	fmt.Println(string(fd2), "\n")

	fmt.Println("=============================")

	o2, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println("3) Seek Error")
		return
	}

	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8) // offset 위치 부터 읽어옴
	if err != nil {
		fmt.Println("3) ReadAt Error")
		return
	}

	fmt.Printf("읽기 작업 완료 (3) (%d bytes) (%d ret)\n\n", ct3, o2)
	fmt.Println(string(fd3), "\n")

	defer file.Close()
}
