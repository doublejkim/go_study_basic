package section10

import (
	"errors"
	"fmt"
	"math"
)

func myPower(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("0은 사용 불가")
	}
	return math.Pow(f, i), nil
}

func ErrorEx1() {

	if f, err := myPower(7, 3); err != nil {
		fmt.Printf("Error Message : %s", err)
	} else {
		fmt.Println("myPower result 1 : ", f)
	}

	if f, err := myPower(0, 3); err != nil {
		fmt.Printf("Error Message : %s", err)
	} else {
		fmt.Println("myPower result 2 : ", f)
	}

}
