package section10

import (
	"fmt"
	"log"
	"math"
	"time"
)

type PowError struct {
	time    time.Time // 에러 발생 시간
	value   float64   // 파라미터 : interface{} 타입으로 하여 어떤 값이라도 받을수있도록하면 활용에 좋음
	message string    // 에러 메시지
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - InputValue(value: %g) - %s", e.time, e.value, e.message)
}

func Power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "0은 사용할 수 없음 "}
	}
	if math.IsNaN(f) {
		return 0, PowError{time: time.Now(), value: f, message: "숫자가 아님"}
	}
	if math.IsNaN(i) {
		return 0, PowError{time: time.Now(), value: f, message: "숫자가 아님"}
	}
	return math.Pow(f, i), nil
}

func ErrorEx2() {

	v, err := Power(10, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result 1 : ", v)

	t, err := Power(10, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result 2 : ", t)
}
