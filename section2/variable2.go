package section2

import "fmt"

func Variable2() {

	var (
		name      string = "machine"
		height    int32
		weight    float32
		isRunning bool
	)

	name = "new name"
	height = 180
	weight = 70.5
	isRunning = true

	fmt.Println("name : ", name)
	fmt.Println("height : ", height)
	fmt.Println("weight : ", weight)
	fmt.Println("isRunning : ", isRunning)

}
