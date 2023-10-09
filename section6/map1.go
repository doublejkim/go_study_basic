package section6

import "fmt"

func Map1() {

	fmt.Println("======================= Map 기본 선언 방법")

	var map1 map[string]int = make(map[string]int)
	var map2 = make(map[string]int)
	map3 := make(map[string]int)

	fmt.Println("map1 : ", map1)
	fmt.Println("map2 : ", map2)
	fmt.Println("map3 : ", map3)

	fmt.Println("======================= Map 기본 선언 방법")

	map4 := map[string]int{}
	map4["apple"] = 10
	map4["banana"] = 5
	map4["cherry"] = 23

	map5 := map[string]int{
		"apple":  21,
		"banana": 31,
		"cherry": 41,
	}

	map6 := make(map[string]int, 10) // 10 개를 make 함수로 초기화해서 할당

	map6["art"] = 51
	map6["brigde"] = 61
	map6["cap"] = 71

	fmt.Println("map4 : ", map4)
	fmt.Println("map5 : ", map5)
	fmt.Println("map6 : ", map6)
	fmt.Println("art : ", map6["art"], ", cap : ", map6["cap"], ", unknown : ", map6["xxx"])

}
