package section3

import (
	"fmt"
	"math/rand"
	"time"
)

func Switch2() {

	//rand.Seed()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch i := r.Int31n(100); {
	case i > 50 && i < 100:
		fmt.Println("i -> ", i, " 50 이상 100미만")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, " 25 이상 50미만")
	default:
		fmt.Println("i -> ", i, " 기본 값")
	}
}
