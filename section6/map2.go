package section6

import "fmt"

func Map2() {

	fmt.Println("======================= Map 에서 값 획득")

	map1 := map[string]string{
		"daum":   "http://www.daum.net",
		"naver":  "http://www.naver.com",
		"google": "http://www.google.com",
	}

	fmt.Println("daum : ", map1["daum"])
	fmt.Println("naver : ", map1["naver"])
	fmt.Println("google : ", map1["google"])

	fmt.Println("======================= Map 순회")

	for k, v := range map1 {
		fmt.Println("Key : ", k, ", Value : ", v)
	}
}
