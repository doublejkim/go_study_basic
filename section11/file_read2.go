package section11

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func FileRead2() {

	// 파일 읽기
	// csv 파일 읽기 예제
	// 패키지 저장소를 통해서 엑셀등 다양한 파일 형식 쓰기, 읽기 가능
	// 패키지 github 주소 : https://github.com/tealeg/xlsx
	// bufio : 파일 용량이 클 경우 버퍼 사용 권장

	// 파일 생성
	file, err := os.Open("files/sample.csv")
	if err != nil {
		fmt.Println(" 1) Open Error")
		return
	}

	defer file.Close()

	// csv reader 생성
	//rr := csv.NewReader() // 몇개 안될때는 그냥 읽어와도 무관..
	rr := csv.NewReader(bufio.NewReader(file)) // 권장

	// csv 내용 읽기(1)
	row1, err1 := rr.Read() // 1개의 row 단위로 읽기
	if err1 != nil {
		fmt.Println(" 2-1) Read Error")
		return
	}

	row2, err2 := rr.Read()
	if err2 != nil {
		fmt.Println("2-2) Read Error")
		return
	}

	fmt.Println("CSV Read Example")
	//fmt.Println(row1)
	//fmt.Println(row1[0]) // slice 로 가져오기때문에 그중 0번째 항목을 획득 가능

	fmt.Println(row1[0], row1[1], row1[7], row1[1:5])
	fmt.Println(row2[0], row2[1], row2[7], row2[1:5])

	fmt.Println("==================================")

	// csv 내용 읽기(2)
	rows, err := rr.ReadAll()
	if err != nil {
		fmt.Println("3) ReadAll Error")
		return
	}

	fmt.Println("CSV ReadAll Example")
	fmt.Println(rows)
	fmt.Println("==================================")

	// row 단위로 csv 파일 읽기
	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s ", rows[i][j])
		}
		fmt.Println()
	}

}
