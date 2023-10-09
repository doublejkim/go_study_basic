package section6

import "fmt"

func Map4() {

	fmt.Println("======================= Map 의 값 획득시 존재 여부 확인")

	map1 := map[string]int{
		"apple":  10,
		"banana": 110,
		"cherry": 200,
		"grape":  20,
	}

	value1, exist1 := map1["apple"]
	value2 := map1["kiwi"]

	// 참고 : map 의 값을 받을때 2개의 값을 받으면 첫번째 값은 value 데이터 타입의 값 (없다면 기본 초기화값)
	//        두번째 값은 실제로 해당 키에 매핑되는 값이 존재하는지 여부를 리턴 해줌
	value3, exist3 := map1["kiwi"]

	fmt.Println("value1 : ", value1, ", exist1 : ", exist1)
	fmt.Println("value2 : ", value2)
	fmt.Println("value3 : ", value3, ", exist3 : ", exist3)

	fmt.Println("======================= Map 값 획득시 존재여부를 조건문으로 체크후 활용")

	if value, ok := map1["kiwi"]; ok {
		fmt.Println("OK. value : ", value)
	} else {
		fmt.Println("Kiwi is not exists...")
	}

	fmt.Println("======================= Map 에서 특정키의 존재여부만 확인하여 활용")

	if _, ok := map1["lemon"]; !ok {
		fmt.Println("lemon is not exists...")
	}
}
