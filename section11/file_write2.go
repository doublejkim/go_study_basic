package section11

import (
	"encoding/csv"
	"fmt"
	"os"
)

func FileWrite2() {

	file, err := os.Create("test_write.csv")
	if err != nil {
		fmt.Println("file open error")
		return
	}

	defer file.Close()

	wr := csv.NewWriter(file)
	// wr := csv.NewWriter(bufio.NewRiter(file)) // 버퍼 사용을 더 권장

	wr.WriteAll([][]string{{"Kim", "M", "99"}, {"Lee", "W", "100"}, {"Park", "M", "70"}})
	wr.Flush()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat() error : ", err)
		return
	}

	fmt.Printf("CSV 쓰기 작업 후 파일 크기 (%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일명 : ", fi.Name())
	fmt.Println("운영체제 파일 권한 : ", fi.Mode())
}
