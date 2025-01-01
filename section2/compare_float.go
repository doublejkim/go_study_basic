package section2

import (
	"fmt"
	"math"
	"math/big"
)

func CompareFloatMistake() {

	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v", a, b, c, a+b == c) // false

}

func CompareFloat1() {

	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v", a, b, c, EqualWithEpsilon(a+b, c)) // true
}

const epsilon = 0.000001

func EqualWithEpsilon(a, b float64) bool {
	// epsilon 만큼의 작은 오차 무시하기
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

func CompareFloat2() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v", a, b, c, EqualWithNextafter(a+b, c)) // true
}

func EqualWithNextafter(a, b float64) bool {

	// math.Nextafter 는 b 를 향해 1비트만큼 더하거나 빼는 메소드
	// epsilon 으로 임의 계산하는것보다는 정밀도를 높혀서 정확한 수치를 계산 가능
	// 단, 어디까지나 오차를 무시하는 방법임
	return math.Nextafter(a, b) == b
}

func CompareFloat3() {
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d))
}
