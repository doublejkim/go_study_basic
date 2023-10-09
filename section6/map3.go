package section6

import "fmt"

func Map3() {
	fmt.Println("======================= Map 값 변경 및 삭제")

	map1 := map[string]string{
		"daum":   "http://www.daum.net",
		"naver":  "http://www.naver.com",
		"google": "http://www.google.com",
		"home1":  "http://test1.com",
	}

	fmt.Println("map1-1 : ", map1)

	map1["home2"] = "http://test2.com" // 기존에 없었던 키의 경우 추가 됨

	fmt.Println("map1-2 : ", map1)

	map1["home1"] = "Unknown" // 특정 키에 대한 데이터 수정

	fmt.Println("map1-3 : ", map1)

	delete(map1, "home1") // 특정 키에 대한 데이터 삭제

	fmt.Println("map1-4 : ", map1)
}
