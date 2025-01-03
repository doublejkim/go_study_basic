# Go Study - Basic

---

# Section 1

## 1.1. Hello World

[helloworld.go](section1/helloworld.go)


# Section 2

## 2.1. 변수  

### 2.1.1. 변수 선언 및 사용

- `var` 로 선언 및 사용
- 값을 초기화 하지않으면 데이터 타입을 같이 정의 해야함 
- 값을 초기화 하면서 선언할 경우는 초기화 값을 보고 데이터타입을 추론하기에 사용 가능 

### 2.1.2. 기본 초기화

1. 정수 타입 : `0`
2. 실수 (소수점) : `0.0`
3. 문자열 : `""`
4. Boolean : `true`, `false`

### 2.1.2. 변수명 선언 규칙

- 숫자로 시작할 수 없음
- 대소문자 구분 
- 문자, 숫자, 밑줄, 특수기호 

변수 초기화 및 선언 예제 : [variable1.go](section2/variable1.go)

### 2.1.3. 변수 그룹화 선언

```go
var (
name      string = "machine"
height    int32
weight    float32
isRunning bool
)
```

변수 그룹화 선언 : [variable2.go](section2/variable2.go)

### 2.1.4. 변수 짧은 선언

- 함수 안 에서만 사용 가능
- 전역으로는 사용 불가
- 선언 후 할당 하면 예외 발생
- var 와 데이터 타입을 생략하고 사용 가능 
- 반복문, 조건문의 초기화 등에서 사용 가능 

```go
shortVar := 5
```
짧은 선언 예제 : [variable3.go](section2/variable3.go)

## 2.2. 실수의 비교 (실수의 오차)

### 2.2.1. 일반적인 실수의 비교시 발생하는 버그

```go
func CompareFloatMistake() {

	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v", a, b, c, a+b == c) // false

}
```

- 실수의 표현방식 때문의 위와 같은 오차 발생 함

### 2.2.2. 작은 오차 무시 

```go
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
```

- 작은 오차는 무시하도록 custome epsilon 비교 메소드를 만들어 활용 가능 

```go
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
```

- `math` 패키지의 `func Nextafter(x, y float64) (r float64)` 활용하면 조금 더 정확한 정밀하게 오차 무시가능
- 해당 함수는 `x` 값을 `y` 값을 향해 1비트만큼 더하거나 빼서 리턴하는 함수. 즉, 해당 데이터타입에서 허용하는 가장 작은 차이로 `+/-` 해서 오차 Skip 하여 비교 가능
- 단, 어디까지나 두가지경우 모두 오차를 무시하는 방법임 

### 2.2.3. math/big 패키지를 이용한 실수(float) 연산 

```go
func CompareFloat3() {
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d)) // true
}
```

- `math/big` 에서 제공하는  `big.Float` 의 `(x *Float) Cmp(y *Float) int` 으로 정밀한 비교가능
  - x 가 작을경우 : -1 / x 가 큰 경우 : 1 / x, y 가 동일한경우 : 0
- `big.Float` 으로 생성한 실수끼리의 연산은 위의 `Add()` 처럼 별도로 함수를 제공

## 2.3. 상수

- 상수 선언 및 사용 : `const`

### 2.3.1. 상수의 사용 

- `const` 키워드를 사용하여 쵸기화
- 한 번 선언 후 값 변경 금지
- 고정 된 관리 용

```go
const a int32 = 10

a = 1 // 에러 발생 
```

상수의 사용 : [const1.go](section2/const1.go)

### 2.3.2. 여러개의 상수 선언

- 동일한 데이터타입의 상수를 한 번에 선언 가능
- 다른데이터타입의 상수도 한 번에 선언 가능
- 변수 선언과 마찬가지로 그룹화해서 선언 가능

여러개의 상수 선언 예제 : [const2.go](section2/const2.go)

## 2.4. 열거형 (Enumeration)

### 2.4.1. 열거형 사용 

- 상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음 
- `iota` 키워드를 사용하면 열거형 사용가능 
- 값의 변형이 필요할경우 `iota` 에 덧셈, 곱셈, 시프트 연산등을 통해 변형하여 사용가능
- `iota` 의 증가값의 중간 생략이 필요할 경우 밑줄(`_`)로 생략가능

열거형 사용 예제 : [enumeration1.go](section2/enumeration1.go)


# Section 3 - 제어문, 반복문 

## 3.1. if 문 

- 반드시 bool 검사 필요
- 괄호 사용. 단 1개 라인만 기술되더라도 괄호 필요 

if 문 사용 예제 : [if1.go](section3/if1.go)

```go
func If1() {

	var a int = 15
	b := 15

	if a >= 15 {
		fmt.Println("a 는 15 이상")
	}

	if b >= 30 {
		fmt.Println("b 는 30 이상")
	}

	if a >= 0 && a < 10 {
		fmt.Println("a >= 0 && a <10")
	} else if a >= 10 && a < 20 {
		fmt.Println("a >= 10 && a < 20 ")
	} else {
		fmt.Println("Unknown....")
	}

}
```

## 3.2. switch 문 

- switch 뒤 표현식(Expression) 생략 가능
- case 뒤 표현식(Expression) 사용 가능
- 자동 break 로 인한 fallthrough 존재
- Type 으로 분기 가능 -> 값이 아닌 변수 Type 으로 분기 가능

switch 문 사용 예제 1 : [switch1.go](section3/switch1.go) <br>
switch 문 사용 예제 2 : [switch2.go](section3/switch2.go) <br>
switch 문 사용 예제 3 : [switch3.go](section3/switch3.go) <br>

```go
func MyFunc() {

	a := 30 / 15
	switch a {
	case 2, 4, 6, 8:
		fmt.Println("2, 4, 6, 8")
	case 1, 3, 5, 7, 9:
		fmt.Println("1, 3, 5, 7, 9")
	default:
		fmt.Println("Default!!!!")
	}

	switch e := "go"; e {
	case "java":
		fmt.Println("Java!!!")
		fallthrough
	case "go":
		fmt.Println("Go!!!")
		fallthrough
	case "python":
		fmt.Println("Python!!!")
		fallthrough
	case "ruby":
		fmt.Println("Ruby!!!")
		fallthrough
	case "c":
		fmt.Println("C!!!")
	}
}
```

## 3.3. for 문

### 3.3.1. 많이 사용되는 일반적인 for 문 

- Golang 에서 유일한 반복문
- 일반적으로 알고있는 for 의 구조와 동일
- 배열의 경우 range 순회 기능 존재 (인덱스와 값을 동시에 받아 사용 가능)
- 인덱스와 값을 동시에 획득시 밑줄(`_`) 을 사용하여 특정값을 생략할 수 있음

```go
// 일반적인 for 
for {초기화 문}; {조건 판별}; {증감} {
	// ....
}

// 배열을 받아 range 로 순회 
loc := []string{"AA", "BB", "CC"}

for index, name := range loc {
    fmt.Println("index : ", index, ", name : ", name)
}
```

for 사용 예제 1 : [for1.go](section3/for1.go)

### 3.3.2. 다양한 형태의 for 문 

- 조건문만 표시해서 while 문 처럼 사용가능
- 초기화문, 조건판별, 증감식에서 사용되는 변수를 여러개 사용 가능 
- `break` 로 반복문 중지 가능

```go
// 조건문만 for 에 표시하여 while 처럼 사용 가능
sum, i := 0, 0
for i <= 100 {
	sum += i
	i ++
}
```

for 사용 예제 2 : [for2.go](section3/for2.go)

### 3.3.3. break, continue 와 label 을 활용한 for 문 제어 

- `break` 사용시 label 명 을 기술하여 원하는 for 문을 중지 시킬 수 있음
- `break` 만 사용시에는 가장 가까운 for 문만 중지 됨 

for 사용 예제 3 : [for3.go](section3/for3.go)

## 3.3. golang 의 특징 

- 후위 연산자 허용 (`i++`), 전위 연산자 비허용 (`++i`)
- 증감 연산자는 반환 값 없음
- 포인터(Pointer) 허용. 단 연산은 비허용
- 문장 끝에 세미콜론 없음.
- 두 문장을 한 문장에 표현할 경우 명시적으로 세미콜론 사용 (단, 권장하지는 않음)
- 조건문, 반복문 등 괄호를 적을때는 개행 후 적으면 에러 발생

# Section 4. 패키지 및 접근제어

## 4.1. 패키지 사용 방법

```go
import "fmt"  // fmt 패키지를 사용 import

// 여러 패키지를 사용 하도록 import
import (
	"fmt"
	"os"
)
```

- 패키지 이름 = 디렉토리 이름
- 같은 패키지 내 -> 소스파일들은 디렉토리명을 패키지명으로 사용
- 네이밍 규칙 : 소문자 private, 대문자 public
- main 패키지는 특별하게 인식 ->  프로그램의 시작점으로 인식 

## 4.2. 접근 제어 및 Alias

- 변수, 상수, 함수, 메서드, 구조체등 식별자의 접근 제어
  - 대문자 : 패키지 외부에서 접근 가능 
  - 소문자 : 패키지 외부 접근 불가 (패키지 내 에서만 접근 가능)

```go
import (
	"fmt"
	mylib "aaa/complexlibrary" // mylib 이라는 alias 사용 가능
	
	// complexlibrary.SomeFunction() 형태가 아닌
	mylib.SomeFunction() // 형태로 alias 사용 가능 
)
```

- 특정 패키지를 사용하지 않더라도 임시로 import 구문을 기술하여놓을 필요가 있을경우 밑줄을 앞에다 적어서 임시 사용가능
- 패키지를 사용하지 않은경우 보통 문법체크하여 import 가 삭제하는데 밑줄(빈 식별자) 표시를 한 경우는 삭제 되지 않음

```go
import (
	"fmt"
	_ "os"   // 파일 내에 os 패키지 사용이 없더라도 저장시 삭제 되지 않음 
)
```

## 4.3. 초기화 함수 

- `init()` 으로 해당 패키지내의 함수가 호출되기 전 호출
- 초기화가 필요한 경우 해당 함수 활용
- 여러개 사용 가능하지만 권고하지 않음 

# Section 5. Go 데이터 타입 (자료형)

## 5.1. bool

- 참 과 거짓 표현
- 조건부 논리 연산자와 주료 사용 : `!`, `&&`, `||`
- 암묵적 형변환 불가(X) : 0, null, Nil -> false 변환 없음

bool 타입 예제 1 : [bool1.go](section5/bool1.go)

## 5.2. 숫자형 기초 (Numeric types)

- 정수, 실수, 복소수
- 32bit, 64bit, unsigned
- 정수 : 8진수(0), 16진수(0x), 10진수 
> [Golang 스펙 - Numeric types](https://go.dev/ref/spec#Numeric_types)

숫자형 예제 1 : [numeric1.go](section5/numeric1.go)

- 실수(부동소수점)
- float32(소수7자리), float64(소수15자리)

숫자형 예제 2 : [numeric2.go](section5/numeric2.go)

## 5.3. 숫자형 연산

- 타입이 같아야 연산 가능
- 다른 타입끼리는 반드시 형 변환 후 연산
- 형 변환 없을 경우 예외(에러) 발생

> [Golang 스펙 - Operators](https://go.dev/ref/spec#Operators)

숫자형 연산 예제 1 : [number_operation1.go](section5/number_operation1.go)

## 5.4. 문자열 기초 

- 큰 따옴표 (`"`), 백스쿼트 로 표현
- 문자 char 타입 존재 하지 않음 -> 대신 rune(int32) 사용. 문자 코드값
- 문자 : 작은 따옴표 (`'`) 로 표현 
- 자주 사용하는 escape : `\\`, `\'`, `\"`, `\a`, `\b`, `\n`, `\r`, `\t`

> [Golang 스펙 - String types](https://go.dev/ref/spec#String_types)

문자열 기초 예제 1 : [string1.go](section5/string1.go)

- 문자열 : UTF-8 인코딩 (유니코드 문자 집합
- 배열 인덱스로 접근할수있지만, 실제 언어별 한글자를 의미하는 한 요소에 접근하고 싶다면 `[]rune` 타입으로 캐스팅후 루핑하면 가능

문자열 기초 예제 1 : [string2.go](section5/string2.go)

## 5.5. 문자열 연산 

### 5.5.1. 문자열 추출 

- 슬라이싱을 사용하면 문자열을 추출 가능 (like substring)
- 1개만 추출하면 문자 번호를 추출

문자열 예제 1 : [string_oper1.go](section5/string_oper1.go)

### 5.5.2. 문자열 비교 

- 같음(`==`), 같지 않음(`!=`) 비교는 가능
- 크기비교는 아스키코드값 기준으로 앞에서부터 비교함 (`<, >`)

문자열 예제 2 : [string_oper2.go](section5/string_oper2.go)

### 5.5.3. 문자열 결합

- plus(`+`) 연산을 통해 단순 결합 가능 
- join 을 통하여 성능을 고려하여 결합 가능 

문자열 예제 3 : [string_oper3.go](section5/string_oper3.go)

# Section 6. 데이터타입 Advanced (array, slice, map, pointer)

## 6.1. 배열 (Array)

### 6.1.1. 배열 선언 및 사용

- 배열은 용량, 길이 항상 같음 
- 배열 : 길이 고정, 값 타입, 복사 전달, 비교 연산자 사용 가능
- 슬라이스 : 길이 가변, 참조 타입, 참조 값 전달, 비교 연산자 사용 불가능 
- `cap()` : 배열, 슬라이스 용량 구하는 함수
- `len()` : 배열, 슬라이스 갯수 구하는 함수

배열 사용 예제 : [arra1.go](section6/array1.go)

### 6.1.2. 배열 순회 

- 배열의 길이 만큼 순회하는 전통적인 방식으로 순회 가능
- `range` 키워드를 활용한 index, value 를 받아 순회 가능 

```go
// 전통적 방식의 순회
for i := 0; i < len(arr); i++ {
	//...
}

// range 활용 순회 
for i, v := range arr {
	//...
}
```

배열 순회 예제 : [array2.go](section6/array2.go)

### 6.1.3. 배열의 값 복사

- 배열을 복사하면 값 복사로 복사 됨

배열 값 복사 예제 : [array3.go](section6/array3.go)

## 6.2. 슬라이스 기초

### 6.2.1. 슬라이스 선언 및 사용

- 슬라이스의 길이는 가변 
- 동적으로 크기가 늘어남 
- 레퍼런스(참조 값 타입)
- 선언 방법 :
  1. 배열처럼 선언 (길이는 명시하지 않고 사용)
  2. `make` 함수 (자료형, 길이, 용량(생략시 길이))

슬라이스 기초 예제 1 : [slice_basic1.go](section6/slice_basic1.go)

### 6.2.2. 슬라이스 참조 복사 / 순회 

- 슬라이스를 복사하면 참조값이 복사 됨
- 메모리상의 동일한 슬라이스를 참조하는 것임
- 주의 : `make` 로 슬라이스 선언 시 길이 만큼만 초기화되기때문에 capa 를 별도로 지정한 경우는 panic 을 조심하자

슬라이스 기초 예제 2 : [slice_basic2.go](section6/slice_basic2.go)

## 6.3. 슬라이스 심화

### 6.3.1. 슬라이스 추가 및 병합 

- 슬라이스에 요소나 다른 슬라이스를 추가하려면 `append` 사용
- 슬라이스의 capa 는 꽉 찼을경우 2배씩 증가 됨 

슬라이스 심화 예제 1 : [slice_ex1.go](section6/slice_ex1.go)

### 6.3.2. 슬라이스 추출 및 정렬

- 슬라이싱 연산으로 subtring 가능 

```go
slice[i:j] // i ~ j-1 까지 추출 
slice[i:]  // i ~ 마지막까지 추출 
slice[:j]  // 처음 ~ j-1 까지 추출 
slice[:]   // 처음 ~ 마지막 까지 추출 
```

- 슬라이스 정렬시 기본 자료형은 `sort` 패키지내에 오름차순정렬하는 기본 함수 제공
- 그 외에 정렬을하거나 내림차순이 정렬이 필요하면 [데이터타입]Slice 형태의 타입을 제공하는데, 해당 타입으로 캐스팅후 Sort 함수를 사용

```go
// 기본자료형은 오름차순 함수를 기본으로 제공함 
sort.Ints(slice)
sort.Float64s(slice)
sort.Strings(slice)   

// 내림 차순 정렬이 필요하면 아래와 같이 사용
sort.Sort(sort.Revers(sort.IntSlice(slice)))
...
```

슬라이스 심화 예제 2 : [slice_ex2.go](section6/slice_ex2.go)

### 6.3.3. 슬라이스 값 복사 

- `copy` 로 복사 가능
- `make` 로 공간을 할당 후 복사 해야함. 공간이 할당되지 않으면 복사되지 않음
- 복사 된 슬라이스 의 값을 변경해도 원본에는 영향 없음 
- 주의 : 부분적인 슬라이스 추출은 참조 가 되어버림
- 부분 슬라이스 추출할때 3번째 값을 적어주면 참조되어지는 Target 슬라이스의 capa 를 지정할 수 있음 

슬라이스 심화 예제 3 : [slice_ex3.go](section6/slice_ex3.go)

## 6.4. 맵 (Map)

### 6.4.1. 맵 - 선언과 사용

- Key-Value 형태로 자료 저장
- 레퍼런스 타입 (참조타입) 
- 비교 연산자 사용 불가능 (레퍼런스 타입 이므로)
- make 함수 및 축약(리터럴) 초기화 가능
- 순서 보장되지 않음 
- 특징
  1. Key 를 사용할때는 레퍼런스 타입은 사용 불가능
  2. Value 는 모든 타입 사용 가능 

> [Golang 스펙 - Map types](https://go.dev/ref/spec#Map_types)
 
맵의 선언과 사용 예제 : [map1.go](section6/map1.go)

### 6.4.2. 맵 - 순회 

- `range` 키워드를 활용한 `for` 문으로 순회 가능
- 맵은 순회할때 순서를 보장하지 않음

맵의 순회 예제 : [map2.go](section6/map2.go)

### 6.4.3. 맵 - 값 추가 / 수정 / 삭제

- 추가 : 기존에 없던 키에 데이터를 대입하면 Map 데이터 추가 
- 수정 : 기존에 존재하는 키에 신규 데이터를 대입하면 Map 데이터 수정 
- 삭제 : `delete` 함수를 이용하여 삭제 시킬 Map 과 Key 를 입력하면 데이터 삭제

```go
myMap := map[string]int {
	"aaa": 10
	"bbb": 20
}

myMap["ccc"] = 30 // 데이터 추가
myMap["bbb"] = 77 // 데이터 수정 
delete(myMap, "bbb") // 데이터 삭제 
```

맵 값 추가, 수정, 삭제 예제 : [map3.go](section6/map3.go)

### 6.4.4. Map 조회시 키가 존재하는지 여부 확인 

- 맵에서 특정 키값에 대한 값을 획득시 존재 여부를 체크하는 값도 같이 받을 수 있음

```go
myMap := map[string]int {
	"aaa": 10,
	"bbb": 20,
}

if value, ok := myMap["ccc"]; ok {
	//...
} else {
	// ...
}

```

맵 조회시 키 존재 여부 확인 예제 : [map4.go](section6/map4.go)

## 6.5. 포인터 (Pointer)

### 6.5.1. 포인터

- Golang : 포인터 지원 
- 주소의 값은 직접 변경 불가능
- 변수 선언시에 asterisk(`*`)를 사용하면 포인터 변수를 의미
- `nil` 로 초기화

> [Golang 스펙 - Pointer types](https://go.dev/ref/spec#Pointer_types)

```go
var a *int // nil 로 초기화
var b *int = new(int) // 임의의 메모리 공간이 할당 됨 
```

- ampersand(`&`) 는 변수의 주소를 가리킴
- 이미 선언되어서 사용되고있는 변수에 asterisk(`*`) 를 사용하면 역참조

```go
i := 7

p1 := &i
p2 := &i

fmt.Println("i : ", i, ", &i : ", &i) // 7, 0x0100
fmt.Println("*p1 : ", *p1, ", p1 : ", p1) // 7, 0x0100
fmt.Println("*p1 : ", *p2, ", p1 : ", p2) // 7, 0x0100

*p2 = 123 // p2 포인터 변수가 레퍼런스하고있는 메모리주소에 저장된 값을 변경

fmt.Println("i : ", i, ", &i : ", &i) // 123, 0x0100
fmt.Println("*p1 : ", *p1, ", p1 : ", p1) // 123, 0x0100
fmt.Println("*p1 : ", *p2, ", p1 : ", p2) // 123, 0x0100
```

포인터 선언 및 참조 사용 예제 : [pointer1.go](section6/pointer1.go)

# Section 7. 함수, Defer, Closure

## 7.1. 함수 Basic

### 7.1.1. 함수 - 선언

- `func` 키워드로 선언
- 여러 개의 반환 값(return value) 사용 가능 

```go
func 함수명(매개변수) (반환타입 or 반환 값 변수명) : 반환 값 존재 
func 함수명() (반환타입 or 반환 값 변수명) : 매개변수 없음, 반환 값 존재 
func 함수명() : 매개변수 없음, 반환 값 없음
```

> [Golang 스펙 - Function types](https://go.dev/ref/spec#Function_types)

### 7.1.2. 함수 - 선언 및 매개변수, 리턴값 사용

- 매개변수와 리턴값 사용 
```go
func myFunc() {
	result := sum(3, 5)  // result : 8
}

func sum(a int, b int) int {
    return a + b
}
```

- 리턴값 여러개 사용
- 가변인자 사용 가능 
- 특정 리턴값을 생략하고 싶으면 밑줄(`_`) 로 생략 가능
```go
func myFunc() {

	result1, result2 := PlusMinusTest(5) // result1 : 6, result2 : 4 
	result3 := SumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) // result3 : 55
    slice1 := []int{5, 6, 7, 8}
    result4 := SumAll(slice1...) // result4 : 26
}

func PlusMinusTest(num int) (result1 int, result2 int) {
	result1 = num + 1
	result2 = num - 1
	return
}

func SumAll(num ...int) int {
	total := 0
	for _, value := range num {
		total += value
	}
	return total
}
```

함수 선언 및 사용 예제 : [func1.go](section7/func1.go)

## 7.2. 함수 Advanced

### 7.2.1. 함수 - 변수에 저장 (like function pointer)

- C 언어의 함수 포인터처럼 함수명으로 레퍼런스가 가능함
- 함수를 레퍼런스 할 수 있는 변수를 사용가능하므로 이 변수를 응용 가능
- 함수의 매개변수로 사용하여 callback 으로 활용할수도 있음
- 슬라이스나 맵에 저장하여 활용 가능 
- 함수를 선언하지 않고 익명함수 사용가능 

```go
func myFunc() {
	
	fmt.Println("======================= Callback Test")
	CallBackTest1(11, 22, TestSum)
	
	fmt.Println("======================= 슬라이스에 함수 저장 및 활용")
	sliceFunc := []func(int, int) int{TestSum, TestMinus, TestMultiply, TestDivide}
	for idx, f := range sliceFunc {
		fmt.Println("idx : ", idx, " : ", f(9, 3))
	}
	
	fmt.Println("======================= 맵에 함수 저장 및 활용")
	mapFunc := map[string]func(int, int) int{
		"sum":      TestSum, 
		"multiply": TestMultiply,
	}
	
	fmt.Println("call sum : ", mapFunc["sum"](10, 3))
	fmt.Println("call multiply : ", mapFunc["multiply"](10, 3))
	
	fmt.Println("======================= 익명 함수 사용")
	
	annonymousFunc1 := func(a int) int {
		return a * 2
	}
	
	fmt.Println("annonymousFunc1 : ", annonymousFunc1(30))

    func(msg string) {
        fmt.Println("anonymousFunc2. msg : ", msg)
    }("anonymous test....")

}

func TestSum(a int, b int) int { 
	return a + b
}

func TestMinus(a int, b int) int {
	return a - b
}

func TestMultiply(a int, b int) int {
	return a * b
}

func TestDivide(a int, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func CallBackTest1(num1 int, num2 int, f func(int, int) int) {
	fmt.Println("result call f : ", f(num1, num2))
}
```

함수 선언 및 사용 예제 : [func2.go](section7/func2.go)

### 7.2.2. Defer 함수 (지연 호출 함수)

- 함수가 끝나기 직전에 실행하는 기능
- 다른언어(ex : java)의 finally 구문과 비슷하게 동작
- 주로 리소스 반환, 열린 파일 닫기, Mutex 잠금 해제 등 

```go
// defer 기본적인 사용
func ex_f1() {
  fmt.Println("f1 : start!")
  defer ex_f2()
  fmt.Println("f1 : end!")
}

func ex_f2() {
  fmt.Println("f2 : called!!!!!!")
}

func myFunc() {
  ex_f1()

// 실행결과
// f1 : start!
// f1 : end!
// f2 : called!!!!!!
	
}
```

```go
// defer 사용시 중첩함수를 사용하면 호출순서가 기대와 달라지니 주의
func start(t string) string {
  fmt.Println("start : ", t)
  return t
}

func end(t string) {
  fmt.Println("end : ", t)
}

func deferTest() {
  defer end(start("b")) ///@@@ 
  fmt.Println("in a")
}
// 실행결과
// start :  b
// in a
// end :  b

```

defer 함수 선언 및 사용 예제 : [func3.go](section7/func3.go)


### 7.2.3. Closure

- 함수 안에서 함수를 선언 및 정의할 수 있고, 바깥쪽 함수에 선언된 변수에도 접근할 수 있는 함수 
- 익명함수 사용할 경우 함수를 변수에 할당해서 사용 가능 


```go
multiply := func(x, y int) int {
  return x * y
}

r1 := multiply(5, 10)

fmt.Println("r1 : ", r1)

m, n := 5, 10
sum := func(c int) int {
  return m + n + c
}

r2 := sum(10)
fmt.Println("r2 : ", r2)
```

- 잘못된 사용으로 남발하게 되면 자원 무분별하게 사용 

```go
func myFunc() {
  cntFunc := increaseCnt()
  
  fmt.Println("test ", cntFunc()) // 1
  fmt.Println("test ", cntFunc()) // 2
  fmt.Println("test ", cntFunc()) // 3
  fmt.Println("test ", cntFunc()) // 4
  fmt.Println("test ", cntFunc()) // 5
}

func increaseCnt() func() int {
  n := 0 // 지역변수 (캡쳐됨)
  return func() int {
    n += 1
    return n
  }
}
```

# Section 8. 사용자 정의 타입, 구조체, 인터페이스 
- 상태, 메소드 분리해서 정의 (결합성 없음)
- 사용자 정의 타입 : 구조체, 인터페이스, 기본 타입(int, float, string...), 함수 
- 구조체와 메소드 연결을 통해서 타 언어의 클래스 형식처럼 사용 가능

## 8.1. type 키워드를 사용한 구조체선언, 기본타입 및 함수 재정의 

### 8.1.1. 구조체의 선언 및 기본적인 사용

```go
type Car struct {
	name string
	color string
	price int
	tax int
}

// Car 구조체와 연결한 함수
func (c Car) Price() int {
	return c.name + "'s Price : " + strconv.Itoa(c.price+c.tax)
}

// Car 구조체내부의 값을 변경하려는 목적이려면 주소전달 형식으로 연결 필요 
func (c *Car) SetName(newName string) {
	c.name = newName
}

func myFunc() {
	bmw := Car{name: "520d", price: 50_000_000, color: "white", tax: 5_000_000}
	fmt.Println( bmw.Price()) // 520d's Price : 55000000 
	
	bmw.SetName("아반떼") 
	fmt.Println( bmw.Price()) // 아반떼's Price : 55000000
}
```

구조체의 선언 및 기본적인 사용 예제 : [user_struct1.go](section8/user_struct1.go)

### 8.1.2. 기본타입의 재정의 타입
- **무슨용도로 사용할까?**
```go
type cnt int

func myFunc() {
	var num cnt = 12
	...
}
```

### 8.1.3. 함수의 재정의 타입
```go
type totCost func(int, int) int

func describe(cnt int, price int, fn totCost) {
	fmt.Printf("cnt : %d, price : %d, totalCost : %d", cnt, price, fn(cnt, price))
}

func myFunc() {
	var orderPrice totCost
	orderPrice = func(cnt int, price int) int {
	    return (cnt * price) + 100_000	
    }

    describe(3, 5000, orderPrice) // cnt : 3, price : 5000, totalCost : 115000
}
```

함수의 재정의 타입 사용 예제 : [user_struct2.go](section8/user_struct2.go)

## 8.2. 구조체 (struct)

### 8.2.1. 구조체 기초
- 필드는 갖지만 메서드는 갖지 않음
- 리시버를 통해 메소드와 연결 
- 구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버) 

```go
type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Caculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func myFunc() {
	kim := Account{number: "111-111", balance: 10_000_000, interest: 0.015}
	lee := Account{number: "111-222", balance: 12_000_000}
	
	fmt.Println("kim Calculate : ", int(kim.Caculate())) // kim Calculate :  10150000
	fmt.Println("lee Calculate : ", int(lee.Caculate())) // lee Calculate :  12000000
	
	// 추가적인 선언 방법. 동적할당 
	//선언 방법1 
	var kim2 *Account = new(Account)
	kim2.number = "111-111"
	kim2.balance = 10_000_000
	kim2.interest = 0.015
	
	// 선언 방법2 
	hong2 := &Account{number: "222-222", balance: 15_000_000, interest: 0.04}
	// 선언 방법3 
	lee2 := new(Account)
	lee2.number = "333-333"
	lee2.balance = 130_000_000
	lee2.interest = 0.025
}
```

구조체 선언 및 사용 1 : [struct1.go](section8/struct1.go)
구조체 선언 및 사용 2 : [struct2.go](section8/struct2.go)

- 익명 구조체 사용도 가능 

```go
        // 구조체 익명 선언 
	car1 := struct{ name, color string }{"520d", "red"}
```

익명 구조체 : [struct3.go](section8/struct3.go)

### 8.2.2. reflect 를 활용한 구조체 메타 정보 획득
- Type 획득이후 각 Field 에 접근하여 `Tag`, `Name`, `Type` 등을 확인 가능
- Java reflection 개념과 유사

```go
type Car struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
}

func myFunc() {
	tag := reflect.TypeOf(Car{})
	
	for i := 0; i<tag.NumField(); i++ {
		fmt.Println(tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
    }

	//출력결과
    //차량명 name string
    //색상 color string
    //제조사 company string
	
}
```

### 8.2.3. 중첩 구조체
- 구조체 내에 구조체 타입을 사용 

```go
type Car struct {
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
```

## 8.3. 구조체 Advanced

### 8.3.1. 생성자 패턴 사용

- 구조체 자체에서는 생성자를 지원하지 않음
- 메소드를 만들어 해당 구조체의 포인터타입을 리턴하게 하여 생성자 처럼 사용 가능 

```go
type Account1 struct {
	number   string
	balance  float64
	interest float64
}

// 생성자 패턴
func NewAccount(number string, balance float64, interest float64) *Account {
	return &Account{number, balance, interest}
}

func StructEx1() {

    account1 := NewAccount("111-222", 1_200_000, 0.04)
    fmt.Printf("account1 : %#v\n", account1)
}
```

생성자 패턴 사용 예제 : [struct_ex1.go](section8/struct_ex1.go)

### 8.3.2. 구조체의 메소드 호출 - 값참조, 주소참조

- 값 참조시에는 원래의 값 영향 받지 않음
- 주소 참조시에는 원래의 값 영향 받음 

구조체 메소드 호출 예제 : [struct_ex2.go](section8/struct_ex2.go)

### 8.3.3. 메소드 오버라이딩
- 상속 관계의 구조체에서 상속받는 쪽(Child) 의 메소드 오버라이딩 가능 

```go
type Employee struct {
	name   string
	salary int
	bonus  int
}

func (e Employee) Calculate() int {
	fmt.Println("Employee's Calculate called!!")
	return e.salary + e.bonus
}

func (e Executives) Calculate() int {
	fmt.Println("Executives's Calculate called!!")
	return e.salary + e.bonus + e.specialBonus
}

type Executives struct {
	Employee
	specialBonus int
}

func MyFunc() {
	
	ep1 := Employee{"kim", 2_000_000, 150_000}
	ep2 := Employee{"park", 1_500_000, 200_000}
	
	ex1 := Executives{
		Employee{"lee", 5_000_000, 150_000}, 200_000,
	}
	
	fmt.Println("ep1 : ", ep1.Calculate())
	fmt.Println("ep2 : ", ep2.Calculate())
	
	fmt.Println("ex1 : ", ex1.Calculate()) // Executives 리시버를 사용하는 메소드가 호출됨
	
// 호출 결과
//Employee's Calculate called!!
//ep1 :  2150000
//Employee's Calculate called!!
//ep2 :  1700000
//Executives's Calculate called!!
//ex1 :  5350000


}
```

구조체 메소드 오버라이딩 예제 : [struct_ex3.go](section8/struct_ex3.go)

> 참고 : Golang 에서 객체지향 개념과 대응되는 Keywords <br>
> Encapsulation -> Packages <br>
> Inheritance -> Composition <br>
> Polymorphism -> Interfaces <br>
> Abstraction -> Embedding <br>

## 8.3. 인터페이스 (Interface)

### 8.3.1. 인터페이스 기초 

- 객체의 동작을 표현. 메소드의 집합
- 단순히 동작에 대한 방법만 표시
- 추상화 제공
- 인터페이스의 메소드를 구현한 타입은 인터페이스로 사용 가능
- 아무것도 정의되지 않은 interface 타입은 `nil` 을 리턴함
- interface 에 명시되어있는 구조체의 메소드만 구현하고있다면 구현한것으로 간주하고 interface 를 사용가능 

```go
type Dog struct {
	name   string
	weight int
}

// bite method 구현
func (d Dog) bite() {
	fmt.Println(d.name, " bites!!")
}

type Behavior interface {
	bite()
}

func Interface2() {
	
	fmt.Println("=================== 인터페이스 사용 1")
	dog1 := Dog{"poll", 10}
	var interface1 Behavior
	interface1 = dog1
	interface1.bite()
	
	fmt.Println("=================== 인터페이스 사용 2")
	dog2 := Dog{"marry", 13}
	interface2 := Behavior(dog2)
	interface2.bite()
	
	fmt.Println("=================== 인터페이스 사용 - 배열로 사용해보기 1")
	inters := []Behavior{dog1, dog2}]
    for idx, _ := range inters {
		inters[idx].bite()
	}
	
	fmt.Println("=================== 인터페이스 사용 - 배열로 사용해보기 2")
	for _, val := range inters {
		val.bite()
	}
}
```

인터페이스 사용 예제 1 : [interface1.go](section8/interface1.go) <br>
인터페이스 사용 예제 2 : [interface2.go](section8/interface2.go)

### 8.3.2. 덕 타이핑 (Duck typing)
- 각 값이나 인스턴스의 실제 타입은 상관하지 않고 구현된 메서드로만 타입을 판단하는 방식
- "만약 어떤 새가 오리처럼 걷고, 헤엄치고, 꽥꽥거리는 소리를 낸다면 나는 그 새를 오리라 부르겠다" - Duck test..

```go
type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func (d Dog) bite() {
	fmt.Println(d.name, " - Dog bite!!!")
}

func (d Dog) sounds() {
	fmt.Println(d.name, " - Dog barks!!!")
}

func (d Dog) dogFeature() {
	fmt.Println(d.name, " - dogs are loyal")
}

func (d Cat) bite() {
	fmt.Println(d.name, " - Cat crawls!!!")
}

func (d Cat) sounds() {
	fmt.Println(d.name, " - Cat cries!!!")
}

func (d Cat) catFeature() {
	fmt.Println(d.name, " - cats are independent")
}

type Behavior interface {
	bite()
	sounds()
}

func act(animal Behavior1) {
	animal.bite()
	animal.sounds()
}

func Interface() {

	dog := Dog{"poll", 10}
	cat := Cat{"nabi", 3}

	animals := []Behavior{dog, cat}

	for _, val := range animals {
		act(val)
	}

}

```

인터페이스 사용 예제 3 : [interface3.go](section8/interface3.go)

### 8.3.3. 빈 인터페이스 사용 
- 빈 인터페이스 : 함수 배개변수, 리턴 값, 구조체 필드 등으로 사용 가능 -> 어떤 타입으로도 변환 가능
- 모든 타입을 나타내기 위해 빈 인터페이스 사용

```go
func printContents(v interface{}) {
	fmt.Printf("Type : (%T)\n", v)
	fmt.Println("v : ", v)
}

func myFunc() {

	var a interface{}

	a = 7.5
	printContents(a)

	a = "Golang.."
	printContents(a)

	a = true
	printContents(a)

	a = nil
	printContents(a)

}
```

빈 인터페이스 기본적 사용 예제 : [interface_ex1.go](section8/interface_ex1.go)

### 8.3.4. Type Assertion
- 실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환후 사용 하는 경우
- 인터페이스(타입) -> 형 변환 
- `interfaceVal.(type)` 로 형변환 가능하며 리턴타입 2개 리턴함 `변환된 값(변환이가능하다면..), 변환가능여부(bool)`
- switch 문에서 구분이 가능함 

```go
func InterfaceEx2() {

	fmt.Println("=========== Type Assertion 테스트")

	var a interface{} = 15

	b := a
	c := a.(int)
	//d := a.(float64) // panic : interface conversion

	fmt.Println("a : ", a)
	fmt.Println("reflect.TypeOf(a)", reflect.TypeOf(a))
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	//fmt.Println("d : ", d)

	if v, ok := a.(int); ok {
		fmt.Println("conversion 1 v : ", v)
	}

	if v, ok := a.(float64); ok {
		fmt.Println("conversion 2 v : ", v)
	}
}

func checkType(arg interface{}) {

	switch arg.(type) {
	case bool:
		fmt.Println("This is a bool : ", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int : ", arg)
	case float32, float64:
		fmt.Println("This is a float : ", arg)
	case string:
		fmt.Println("This is a string : ", arg)
	case nil:
		fmt.Println("This is a nil : ", arg)
	default:
		fmt.Println("Unknown type... : ", arg)
	}
}

func InterfaceEx3() {
	checkType(true)
	checkType(12)
	checkType(12.3)
	checkType("what?")
	checkType(nil)
}
```

Type Assertion 예제 : [interface_ex2.go](section8/interface_ex2.go)

# Section 9. Go 병행처리 - 고루틴, 채널, 동기화

## 9.1. 고루틴 (Goroutine)

### 9.1.1. 고루틴 기초
- 타 언어의 스레드(Thread)와 비슷한 기능
- 생성 방법 매우 간단, 리소스 매우 적게 사용 
- 비동기적 함수 루틴 실행 (매우 적은 용량 차지) -> 채널을 통한 통신 가능
- 공유메모리 사용 시에 정확한 동기화 코딩 필요

고루틴 예제 1 : [goroutine1.go](section9/goroutine1.go) <br>
고루틴 예제 2 : [goroutine2.go](section9/goroutine2.go)

```go
// goroutine 예제 1 
func exe1(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Println("exe 1 [", i, "] : ", time.Now())
		time.Sleep(500 * time.Millisecond)
	}
}

func exe2(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Println("exe 2 [", i, "] : ", time.Now())
		time.Sleep(500 * time.Millisecond)
	}
}

func exe3(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Println("exe 3 [", i, "] : ", time.Now())
		time.Sleep(500 * time.Millisecond)
	}
}

func myFunc() {
	cnt := 10
	go exe1(cnt)
	go exe2(cnt)
	exe3(cnt)
}

```

```go
// goroutine 예제 2
func exe(name int) {
	r := rand.Intn(100)
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, " >>>>> ", r, i)
	}
	fmt.Println(name, " end : ", time.Now())
}

func Goroutine2() {

	runtime.GOMAXPROCS(runtime.NumCPU())                        // 현재 시스템의 cpu 코어 개수 반환후 설정
	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(0)) // 설정값 출력

	fmt.Println("Main Routine Start : ", time.Now())
	for i := 0; i < 100; i++ {
		go exe(i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Main Routine End : ", time.Now())
}
```

- Closure 도 goroutine 사용가능
- 반복문안에 있는 클로저는 일반적으로 즉시실행, but 고루틴 클로저는 가장 나중에 실행 (반복문 종료 후 실행)

```go
func myFunc() {

	runtime.GOMAXPROCS(1) // 1개만 사용하겠다고 선언

	s := "Goroutine Closure"

	for i := 0; i < 1000; i++ {
		func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) // 반복문 클로저는 일반적으로 즉시 실행. But, 고루틴 클로저는 가장 나중에 실행 (반복문종료 후)
	}

	time.Sleep(3 * time.Second)
}

```

## 9.2. 채널 (Channel) 기초

### 9.2.1. 채널 

- 고루틴 간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해 사용 : 채널(동기식, 레퍼런스 타입)
- 실행 흐름 제어 가능 (동기, 비동기) -> 일반 변수로 선언후 사용 가능 
- 데이터 전달 자료형 선언 후 사용 가능 (지정 된 타입만 주고 받을 수 있음)
- `interface{}` 전달을 통해서 자료형 상관없 전송 및 수신 가능 
- 포인터 (슬라이스, 맵) 등을 전달시에는 주의 -> 동기화 사용 (Mutex)
- 멀티 프로세싱 처리에서 deadlock 주의  
- `<-` , `->` : `채널 <- 데이터 // 송신` , ` <- 채널 // 수신 `

```go
func work1(v chan int) {
	fmt.Println("Work1 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : E ---> ", time.Now())
	v <- 1
}

func work2(v chan int) {
	fmt.Println("Work2 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : E ---> ", time.Now())
	v <- 2
}

func GoChannel1() {

	fmt.Println("myFunc : S ---> ", time.Now())

	//var c chan int
	//c = make(chan int) // 이렇게 사용해도 문제 없음
	v := make(chan int) // int 형 채널 선언

	go work1(v)
	go work2(v)

	<-v
	<-v
	fmt.Println("myFunc : E ---> ", time.Now())

}
```

채널 기본 사용 예제 : [gochannel1.go](section9/gochannel1.go)

### 9.2.2. 채널로 값 주고 받기 

- 채널을 사용하면 고루틴에서의 값을 받을 수 있음
- 채널로 값을 받을때는 고루틴에서 종료된 순서대로 수신 됨
- 동기식 이라서 채널 수신을 받을때 까지 대기 

```go
func rangeSum(rg int, c chan int) {

	sum := 0

	for i := 0; i <= rg; i++ {
		sum += i
	}
	c <- sum
}

func MyFunc() {

	c := make(chan int)
	go rangeSum(1000, c)
	go rangeSum(7000, c)
	go rangeSum(5000, c)

	// 순서대로 데이터 수신(동기) : 채널에서 값 수신 완료 될 때까지 대기
	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("result1 : ", result1)
	fmt.Println("result2 : ", result2)
	fmt.Println("result3 : ", result3)
}
```

채널로 값 주고 받기 예제 : [gochannel2.go](section9/gochannel2.go)

### 9.2.3. 채널 버퍼 사용 

- 버퍼를 사용하여 비동기식으로 사용가능 
- 버퍼 : 발신 -> 가득차면 대기, 비어있으면 작동, 수신 -> 비어있으면 대기, 가득차있으면 작동 

```go
func MyFunc() {
	// 비동기 : 버퍼사용

	runtime.GOMAXPROCS(1)
	ch := make(chan bool, 2) // 채널 버퍼의 용량을 설정 가능.
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}
}
```

채널 버퍼 사용 예제 1 : [gochannel3.go](section9/gochannel3.go) <br>
채널 버퍼 사용 예제 2 : [gochannel4.go](section9/gochannel4.go)

### 9.2.4. 채널에서 range 로 값 획득 

- 채널을 사용 후 Close 해야하는 것이 일반적 -> 닫힌 채널에 값 전송시 panic 발생 
- Range 를 사용하면 해당 채널이 Close 될때까지 대기하고 값을 꺼내옴 

```go
func MyFunc() {

	// Close : 채널 닫기. 주의 : 닫힌 채널에 값 전송시 panic 발생
	// Range : 채널 안에서 값을 꺼냄 (순회), 채널 닫아야(Close) 반복문 종료 
	//          -> 채널이 열려있고 값 전송하지 않으면 계속 대기

	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch) // 5회 채널에 값 전송 후 채널 닫기
	}()

	for i := range ch { // 채널에서 값을 꺼내옴 (채널 close 될때 까지)
		fmt.Println("i : ", i)
	}
}
```

채널에서 ranage 로 값 획득 예제 : [gochannel5.go](section9/gochannel5.go)

### 9.2.5. 채널에서 값을 가져올 수 있는 상태 확인

- `val, ok` 형태로 채널에서 수신한 값과 채널에서 값을 수신할 수 있는 여부 값을 동시에 리턴 가능
- 조건문에서도 `ok` 값으로 활용 가능 

```go
func MyFunc() {

	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- "Good~!!!"
		}
	}()

	val1, ok1 := <-ch
	fmt.Println("val1 : ", val1, ", ok1 : ", ok1)

	val2, ok2 := <-ch
	fmt.Println("val2 : ", val2, ", ok2 : ", ok2)

	val3, ok3 := <-ch
	fmt.Println("val3 : ", val3, ", ok3 : ", ok3)
	
	// close(ch)
	// _, ok4 := <-ch // ok4 는 false 

}
```

채널에서 값을 가져올 수 있는 상태 확인 예제 : [gochannel6.go](section9/gochannel6.go)

## 9.3. 채널 (Channel) 심화 

### 9.3.1. 채널 : 발신전용, 수신전용

- 함수 등의 매개 변수에 수신 및 발신 전용 채널 지정 가능 
- 전용 채널 설정 후 방향이 다를 경우 예외 발생 (panic)
- 발신 전용 `channel <- 데이터형` 
- 수신 전용 `<- channel`

```go
func sendOnly(c chan<- int, cnt int) {

	for i := 0; i < cnt; i++ {
		c <- i
	}

	c <- 777
	//fmt.Println(<-c) // 에러 발생
}

func receiveOnly(c <-chan int) {
	for i := range c {
		fmt.Println("received : ", i)
	}

	fmt.Println(<-c)
}

func MyFunc() {

	c := make(chan int)

	go sendOnly(c, 10) // 발신 전용
	go receiveOnly(c)  // 수신 전용

	time.Sleep(2 * time.Second)
}
```

발신전용, 수신전용 채널 예제 : [gochannel_ex1.go](section9/gochannel_ex1.go)

### 9.3.2. 발신전용, 수신전용 채널의 함수에서의 활용

- 채널 또한 함수의 반환 값으로 사용 가능

```go
func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 0; i < cnt; i++ {
			sum += i
		}
		tot <- sum
	}()
	return tot
}

func MyFunc() {

	c := sum(100)
	// c<- 200 // error 발생 
	fmt.Println("c : ", <-c)
}
```

함수에서 리턴타입으로 활용하는 수신전용 채널타입 예제 : [gochannel_ex2.go](section9/gochannel_ex2.go), [gochannel_ex3.go](section9/gochannel_ex3.go)

### 9.3.3. 채널 셀렉트 사용 (channel select)

- 여러 채널을 손쉽게 사용할 수 있는 select 분기문 
- 채널에 값이 수신되면 해당 case 문실행 (switch 분기문과 유사함)

```go
func MyFunc() {

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 11
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Hello"
			time.Sleep(700 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("ch1 : ", num)
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
```

## 9.4. 동기화 기초 

### 9.4.1. 동기화를 사용하지 않는 고루틴 

- 동기화를 사용하지 않으면 의도했던 제어가 되지 않음 

```go
type count struct {
	num int
}

func (c *count) increment() {
	c.num += 1
}

func (c *count) result() {
	fmt.Println(c.num)
}

func MyFunc() {

	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i < 10_000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // cpu 양보
		}()
	}

	for i := 1; i < 10_000; i++ {
		<-done
	}

	c.result() // 실행할 마다 결과가 달라짐

}
```

동기화 처리가 되지 않은 동작 예제 : [go_sync1.go](section9/gochannel_ex1.go)

### 9.4.2. 동기화 객체 사용하여 고루틴 사용 (Mutex)

- Mutex 의 `Lock()`, `Unlock()` 을 사용하면 동기화 사용 가능 

```go
type count struct {
	num   int
	mutex sync.Mutex
}

func (c *count) increment() {
	c.mutex.Lock() // 스레드들로부터 연산 보호
	c.num += 1
	c.mutex.Unlock()
}

func (c *count) result() {
	fmt.Println(c.num)
}

func GoSync() {

	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i < 10_000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // cpu 양보
		}()
	}

	for i := 1; i < 10_000; i++ {
		<-done
	}

	c.result()

}
```

Mutex 를 사용한 동기화 예제 : [go_sync2.go](section9/go_sync2.go)

### 9.4.3. 읽기 쓰기 Mutex 사용 
- 특정 로직이 얼마동안 수행될지 모르는 상황에서 특정 값을 공유하며 쓰기/읽기를 한다면 타이밍 이슈 발생

```go
func MyFunc() {

	data := 0

	go func() {
		for i := 1; i <= 10; i++ {
			data += 1
			fmt.Println("Write : data : ", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read 1 : data : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read 2 : data : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
}
```

특정 변수에 동시에 읽기쓰기 예제 : [go_sync3.go](section9/go_sync3.go)


- Read Lock : 읽기 락 끼리는 서로를 막지 않음. 하지만 읽기시도 중에 값이 변경되면 안되므로 쓰기 락은 막음 
- Write Lock : 쓰기 시도 중에 다른 곳에서 이전 값을 읽으면 안 되고, 다른 곳에서 값을 바꾸면 안 되므로 읽기, 쓰기 락 모두 막음 

```go
func MyFunc() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.Lock()
			data += 1
			fmt.Println("Write : data : ", data)
			time.Sleep(200 * time.Millisecond)
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.RLock()
			fmt.Println("Read 1 : data : ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.RLock()
			fmt.Println("Read 2 : data : ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()

	time.Sleep(15 * time.Second)
}
```

쓰기, 읽기 락 사용 예제 : [go_sync4.go](section9/go_sync4.go)

### 9.4.4. Mutext 조건 변수 사용
- 대기하고 있는 객체를 하나 또는 여러 개를 깨울때 사용 
- `NewCond(l Locker)` : 조건 변수 생성 
- `Wait()` : 고루틴 실행을 멈추고 대기
- `Signal()` : 대기하고 있는 고루틴 하나만 깨움
- `Broadcast()` : 대기하고 있는 모든 고루틴을 깨움 

```go
func MyFunc() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) // 비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Waiting... : ", n)
			condition.Wait()
			fmt.Println("Waiting End : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
	}

	/*for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println("Wake Goroutine(Signal) : ", i)
		condition.Signal() // 한개씩 깨움 (모든 고루틴 생성 후)
		mutex.Unlock()
	}*/

	mutex.Lock()
	fmt.Println("Wake Goroutine(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)
}
```

Mutext 조건변수 사용 예제 : [go_sync5.go](section9/go_sync5.go)

## 9.5. 동기화 고급

### 9.5.1. 함수를 한번만 실행 (Once)
- `sync.Once`
- `func(*Once) Do(f func())` : 함수를 한번만 실행 

```go
func onceTest() {
	fmt.Println("Once Test!!!!!!!!!!!!!!") // 한 번만 실행할 코드 
}

func MyFunc() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once)

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println("Goroutine : i : ", n)
			once.Do(onceTest)
		}(i)
	}

	time.Sleep(3 * time.Second)
}
```

동기화 Once 사용 예제 : [go_sync_ex1.go](section9/go_sync_ex1.go)

### 9.5.2. 대기그룹(WaitGroup) 사용 
- 고루틴이 모두 끝날 때 까지 기다릴 때 사용 
- `sync.WaitGroup`
- `func (wg *WaitGroup) Add(delta int)` : 대기 그룹에 고루틴 개수 추가
- `func (wg * WaitGroup) Done()` : 고루틴 끝났다는 것을 알려줄 때 사용
- `func (wg *WaitGroup) Wait()` : 모든 고루틴이 끝날 때까지 기다림 
- Add 로 추가 된 고루틴 개수와 Done 으로 완료된 고루틴 개수가 동일해야 정확하게 동작

```go
func MyFunc() {

	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("WaitGroup : ", n)
			wg.Done()
		}(i)
	}

	// Add == Done 횟수가 같아야함
	wg.Wait()

}
```

대기그룹 사용 예제 : [go_sync_ex2.go](section9/go_sync_ex2.go)

### 9.5.3. 원자적 연산
- 더이상 쪼갤 수 없는 연산 
- 스레드(고루틴), CPU 코어에서 같은변수(메모리)를 수정할때 서로 영향을 받지 않고 안전하게 연산 가능 
- [https://pkg.go.dev/sync/atomic](https://pkg.go.dev/sync/atomic)
- `sync/atomic`
- `Add 계열` : 변수에 값을 더하고 결과를 리턴
- `CompareAndSwap 계열` : 변수 A 와 B 를 비교하여 같으면 C를 대입. 그리고 A,B 같으면 true, 다르면 false
- `Load 계열` : 변수에서 값을 가져옴
- `Store 계열` : 변수에 값을 저장
- `Swap 계열` : 변수에서 새 값을 대입하고, 이전 값을 리턴 

# Section 10. 에러 처리

[Panic 처리 레퍼런스](https://go.dev/ref/spec#Handling_panics)

## 10.1 에러 기초 

### 10.1.1. 에러 (error) 타입
- 에러 발생시 사용하는 타입 
- `Error()` 함수로 에로 내용 확인 가능 

```go
func MyFunc() {

	f, err := os.Open("unnamedFile")

	if err != nil {
		//fmt.Println("error Msg : ", err.Error())
		log.Fatal("error Msg2 : ", err.Error())
	}

	fmt.Println(f.Name())

}
```

에러 타입 사용 예제 : [error1.go](section10/error1.go)

### 10.1.2. 에러 (error) 사용 
- 특정상황에서의 에러를 정의 가능 
- `fmt.Errorf()` 또는 `errors.New()` 로 사용 가능

```go
func notZero(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprintf("Hello Golang : %d", n)
		return s, nil
	}

	return "", fmt.Errorf("%d를 입력했음!!. 에러 발생 : ", n) //errors.New() 로 대체 가능
}

func MyFunc() {

	a, err := notZero(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result a : ", a)

	b, err := notZero(0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result b : ", b)
}
```

```go
func MyFunc() {
	var err1 error = errors.New("Error occurred - 1")
	err2 := errors.New("Error occurred - 2")
	
	fmt.Println("error1 : ", err1)
	fmt.Println("error1 : ", err1.Error())
	
	fmt.Println("error2 : ", err2)
	fmt.Println("error2 : ", err2.Error())
}
```

에러 정의 사용 예제 1 : [error2.go](section10/error2.go) <br>
에러 정의 사용 예제 2 : [error3.go](section10/error3.go)

## 10.2 에러 고급

### 10.2.1. 함수 에서의 리턴으로 활용
- 리턴으로 함수가 기대했던 데로 실행되었는지 여부를 명시할 수 있음 

```go
func myPower(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("0은 사용 불가")
	}
	return math.Pow(f, i), nil
}

func MyFunc() {

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
```

함수에서의 에러 활용 예제 : [error_ex1.go](section10/error_ex1.go)

### 10.2.2. 사용자 에러 정의 

- `error` 인터페이스의 `Error()` 메소드를 구현하여 사용자 에러를 정의 가능

```go
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

func MyFunc() {

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
```

사용자 정의 에러 사용 : [error_ex2](section10/error_ex2.go)

## 10.3. Panic, Recover 의 사용 

### 10.3.1. Panic

- 사용자가 Panic 발생 가능 
- Panic 함수는 호출 즉시, 해당 메소드를 즉시 중지 시키고 defer 함수를 호출하고 자기자신을 호출한 곳으로 리턴 
- 런타임 이외에 사용자가 코드 흐름에 따라 Panic 을 발생시켜 활용할 수 도 있음

```go
func MyFunc() {
	fmt.Println("Start !!!")
	panic("Error occurred : user stopped!!")
	fmt.Println("Finish !!!") // 실행 불가
}
```
panic 사용 예제 : [panic_recover1.go](section10/panic_recover1.go)

### 10.3.2. Recover 

- 에러 복구 가능
- Panic 으로 발생한 에러를 보구 후 코드 재 실행 (프로그램 종료되지 않음 )
- Panic 에서 설정한 메시지를 받아 올 수 있음 

```go
// 에러 발생 recover 예제 1
func runFunc() {
	defer func() {
		s := recover()
		fmt.Println("Error Message : ", s)
	}()
	
	a := [3]int{1, 2, 3}
	
	for i := 0; i < 5; i++ {
		fmt.Println("ex : ", a[i]) // 에러 발생
	}
}

func MyFunc() {

runFunc()
	fmt.Println("Hello???")
}
```

```go
// 에러 발생 recover 예제 2
func fileOpen(filename string) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error : ", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Open File Name : ", f.Name())
	}

	defer f.Close()
}

func MyFunc() {

	fileOpen("test")

	fmt.Println("End Main!!!!")
}
```

recover 예제 1 : [panic_recover3.go](section10/panic_recover3.go) <br>
recover 예제 2 : [panic_recover4.go](section10/panic_recover4.go)

# Section 11. 파일 입출력

## 11.1. 파일 쓰기

### 11.1.1. 파일 쓰기 - 기본
- os 패키지에서 제공 
- `func Create(name string) (file *File, err error)` : 기존 파일을 열거나 새 파일을 생성 
- `func (f *File) Close() error` : 열린 파일을 닫음 
- `func (f *File) Write(b []byte) (n int, err error)` : 파일에 값을 씀. 파일에 쓴 데이터의 길이와 에러 값을 리턴

```go
func MyFunc() {

	file, err := os.Create("test_write1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	str := "Hello World..."

	n, err := file.Write([]byte(str)) // str 을 []byte 슬라이스로 변환. str 을 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, " 바이트 저장 완료")

}
```

파일쓰기 기본 예제 : [file_write1.go](section11/file_write1.go)

### 11.1.2. 파일 쓰기 - CSV
- encoding/csv 패키지를 사용하면 편하게 사용 가능 

```go
func MyFunc() {

	file, err := os.Create("test_write.csv")
	if err != nil {
		fmt.Println("file open error")
		return
	}

	defer file.Close()

	wr := csv.NewWriter(file)
	// wr := csv.NewWriter(bufio.NewRiter(file)) // 버퍼 사용을 더 권장

	wr.WriteAll([][]string{{"Kim", "M", "99"}, {"Lee", "W", "100"}, {"Park", "M", "70"}})
	wr.Flush()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat() error : ", err)
		return
	}

	fmt.Printf("CSV 쓰기 작업 후 파일 크기 (%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일명 : ", fi.Name())
	fmt.Println("운영체제 파일 권한 : ", fi.Mode())
}
```

파일 쓰기 csv 예제 : [file_write2.go](section11/file_write2.go)

## 11.2. 파일 읽기

### 11.2.1. 파일 읽기 기본 
- os 패키지에서 제공
- `func Open(name string) (file *File, err error)` : 파일 읽기
- `func (f *File) Stat() (fi FileInfo, err error)` : 파일의 정보를 얻어옴 
- `func (f *File) Read(b []byte) (n int, err error)` : 파일에서 값을 읽음. 파일에서 읽은 데이터의 길이와 에러 값을 리턴

```go
func FileRead1() {

	fmt.Println("============================= Open/Read 예제")

	file, err := os.Open("files/sample.txt")
	if err != nil {
		fmt.Println(" 1) Open Error")
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("1) Stat Error")
		return
	}

	fd1 := make([]byte, fileInfo.Size()) // 파일 크기 만큼 바이트 슬라이스 생성
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력 : ", fileInfo, "\n")
	fmt.Println("파일 이름 : ", fileInfo.Name(), "\n")
	fmt.Println("파일 크기 : ", fileInfo.Size(), "\n")
	fmt.Println("파일 수정 시간 : ", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업 완료 (1) (%d bytes)\n\n", ct1)
	//fmt.Println(fd1) // byte 값 그대로 출력됨
	fmt.Println(string(fd1))

	fmt.Println("============================= Seek 기본 예제")

	o1, err := file.Seek(20, 0) // whence - 0: 처음 위치, 1: 현재 위치, 2: 마지막 위치
	if err != nil {
		fmt.Println("2) Seek Error")
		return
	}

	fd2 := make([]byte, 20) // 앞에서 offset 20 만큼 이동한상태에서 20만큼 read
	ct2, err := file.Read(fd2)
	if err != nil {
		fmt.Println("2) Read Error")
		return
	}

	fmt.Printf("읽기 작업 완료 (2) (%d bytes) (%d ret)\n\n", ct2, o1)
	fmt.Println(string(fd2), "\n")

	fmt.Println("=============================")

	o2, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println("3) Seek Error")
		return
	}

	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8) // offset 위치 부터 읽어옴
	if err != nil {
		fmt.Println("3) ReadAt Error")
		return
	}

	fmt.Printf("읽기 작업 완료 (3) (%d bytes) (%d ret)\n\n", ct3, o2)
	fmt.Println(string(fd3), "\n")

	defer file.Close()
}
```

파일 읽기 예제 1 : [file_read1.go](section11/file_read1.go)

### 11.2.2. 파일 읽기/쓰기
- `func OpenFile(name string, flag int, perm FileMode) (file *File, err error)` : 파일 플래그 파일 모드를 지정하여 파일 열기
- `func (f *File) Seek(offset int64, whence int) (ret int64, err error)` : 파일을 읽거나 쓸 위치로 이동 
- 참고 : 패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기, 읽기 가능
- 패키지 Github 주소 : https://github.com/tealeg/xlsx
- bufio : 파일이 용량이 클 경우 버퍼 사용 권장 



