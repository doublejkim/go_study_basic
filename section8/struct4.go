package section8

import (
	"fmt"
	"reflect"
)

type Car4 struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
}

func Struct4() {

	tag := reflect.TypeOf(Car4{})

	for i := 0; i < tag.NumField(); i++ {
		fmt.Println(tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}
}

type Car5 struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct {
	length int "전장"
	height int "전고"
	width  int "전축"
}

func Struct5() {

	car := Car5{
		"520d",
		"silver",
		"bmw",
		spec{4_000, 1_000, 2_000},
	}

	fmt.Println("name : ", car.name)
	fmt.Println("color : ", car.color)
	fmt.Println("company : ", car.company)
	fmt.Printf("detail : %#v\n", car.detail)
	fmt.Println("   - length : ", car.detail.length)
	fmt.Println("   - height : ", car.detail.height)
	fmt.Println("   - width : ", car.detail.width)
}
