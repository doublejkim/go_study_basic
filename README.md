# Go Study - Basic

## Section 1

### 1.1. Hello World

[helloworld.go](section1/helloworld.go)


## Section 2

### 2.1. 변수  

#### 2.1.1. 변수 선언 및 사용

- `var` 로 선언 및 사용
- 값을 초기화 하지않으면 데이터 타입을 같이 정의 해야함 
- 값을 초기화 하면서 선언할 경우는 초기화 값을 보고 데이터타입을 추론하기에 사용 가능 

#### 2.1.2. 기본 초기화

1. 정수 타입 : `0`
2. 실수 (소수점) : `0.0`
3. 문자열 : `""`
4. Boolean : `true`, `false`

#### 2.1.2. 변수명 선언 규칙

- 숫자로 시작할 수 없음
- 대소문자 구분 
- 문자, 숫자, 밑줄, 특수기호 

변수 초기화 및 선언 예제 : [variable1.go](section2/variable1.go)

#### 2.1.3. 변수 그룹화 선언

```go
var (
name      string = "machine"
height    int32
weight    float32
isRunning bool
)
```

변수 그룹화 선언 : [variable2.go](section2/variable2.go)

#### 2.1.4. 변수 짧은 선언

- 함수 안 에서만 사용 가능
- 전역으로는 사용 불가
- 선언 후 할당 하면 예외 발생
- var 와 데이터 타입을 생략하고 사용 가능 
- 반복문, 조건문의 초기화 등에서 사용 가능 

```go
shortVar := 5
```

짧은 선언 예제 : [variable3.go](section2/variable3.go)

### 2.2. 상수

- 상수 선언 및 사용 : `const`

#### 2.1.1. 상수의 사용 

- `const` 키워드를 사용하여 쵸기화
- 한 번 선언 후 값 변경 금지
- 고정 된 관리 용

```go
const a int32 = 10

a = 1 // 에러 발생 
```

상수의 사용 : [const1.go](section2/const1.go)

#### 2.1.2. 여러개의 상수 선언

- 동일한 데이터타입의 상수를 한 번에 선언 가능
- 다른데이터타입의 상수도 한 번에 선언 가능
- 변수 선언과 마찬가지로 그룹화해서 선언 가능

여러개의 상수 선언 예제 : [const2.go](section2/const2.go)

### 2.3. 열거형 (Enumeration)

#### 2.3.1. 열거형 사용 

- 상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음 
- `iota` 키워드를 사용하면 열거형 사용가능 
- 값의 변형이 필요할경우 `iota` 에 덧셈, 곱셈, 시프트 연산등을 통해 변형하여 사용가능
- `iota` 의 증가값의 중간 생략이 필요할 경우 밑줄(`_`)로 생략가능

열거형 사용 예제 : [enumeration1.go](section2/enumeration1.go)


## Section 3 - 제어문, 반복문 

### 3.1. if 문 

- 반드시 bool 검사 필요
- 괄호 사용. 단 1개 라인만 기술되더라도 괄호 필요 

if 문 사용 예제 : [if1.go](section3/if1.go)

### 3.2. switch 문 

- switch 뒤 표현식(Expression) 생략 가능
- case 뒤 표현식(Expression) 사용 가능
- 자동 break 로 인한 fallthrough 존재
- Type 으로 분기 가능 -> 값이 아닌 변수 Type 으로 분기 가능

switch 문 사용 예제 1 : [switch1.go](section3/switch1.go) <br>
switch 문 사용 예제 2 : [switch2.go](section3/switch2.go) <br>
switch 문 사용 예제 3 : [switch3.go](section3/switch3.go) <br>

### 3.3. for 문

#### 3.3.1. 많이 사용되는 일반적인 for 문 

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

#### 3.3.2. 다양한 형태의 for 문 

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

#### 3.3.3. break, continue 와 label 을 활용한 for 문 제어 

- `break` 사용시 label 명 을 기술하여 원하는 for 문을 중지 시킬 수 있음
- `break` 만 사용시에는 가장 가까운 for 문만 중지 됨 

for 사용 예제 3 : [for3.go](section3/for3.go)

### 3.3. golang 의 특징 

- 후위 연산자 허용 (`i++`), 전위 연산자 비허용 (`++i`)
- 증감 연산자는 반환 값 없음
- 포인터(Pointer) 허용. 단 연산은 비허용
- 문장 끝에 세미콜론 없음.
- 두 문장을 한 문장에 표현할 경우 명시적으로 세미콜론 사용 (단, 권장하지는 않음)
- 조건문, 반복문 등 괄호를 적을때는 개행 후 적으면 에러 발생

## Section 4. 패키지 및 접근제어

### 4.1. 패키지 사용 방법

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

### 4.2. 접근 제어 및 Alias

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

### 4.3. 초기화 함수 

- `init()` 으로 해당 패키지내의 함수가 호출되기 전 호출
- 초기화가 필요한 경우 해당 함수 활용
- 여러개 사용 가능하지만 권고하지 않음 

## Section 5. Go 데이터 타입 (자료형)

### 5.1. bool

- 참 과 거짓 표현
- 조건부 논리 연산자와 주료 사용 : `!`, `&&`, `||`
- 암묵적 형변환 불가(X) : 0, null, Nil -> false 변환 없음

bool 타입 예제 1 : [bool1.go](section5/bool1.go)

### 5.2. 숫자형 기초 (Numeric types)

- 정수, 실수, 복소수
- 32bit, 64bit, unsigned
- 정수 : 8진수(0), 16진수(0x), 10진수 
> [Golang 스펙 - Numeric types](https://go.dev/ref/spec#Numeric_types)

숫자형 예제 1 : [numeric1.go](section5/numeric1.go)

- 실수(부동소수점)
- float32(소수7자리), float64(소수15자리)

숫자형 예제 2 : [numeric2.go](section5/numeric2.go)

### 5.3. 숫자형 연산

- 타입이 같아야 연산 가능
- 다른 타입끼리는 반드시 형 변환 후 연산
- 형 변환 없을 경우 예외(에러) 발생

> [Golang 스펙 - Operators](https://go.dev/ref/spec#Operators)

숫자형 연산 예제 1 : [number_operation1.go](section5/number_operation1.go)

### 5.4. 문자열 기초 

- 큰 따옴표 (`"`), 백스쿼트 로 표현
- 문자 char 타입 존재 하지 않음 -> 대신 rune(int32) 사용. 문자 코드값
- 문자 : 작은 따옴표 (`'`) 로 표현 
- 자주 사용하는 escape : `\\`, `\'`, `\"`, `\a`, `\b`, `\n`, `\r`, `\t`

> [Golang 스펙 - String types](https://go.dev/ref/spec#String_types)

문자열 기초 예제 1 : [string1.go](section5/string1.go)

- 문자열 : UTF-8 인코딩 (유니코드 문자 집합
- 배열 인덱스로 접근할수있지만, 실제 언어별 한글자를 의미하는 한 요소에 접근하고 싶다면 `[]rune` 타입으로 캐스팅후 루핑하면 가능

문자열 기초 예제 1 : [string2.go](section5/string2.go)

### 5.5. 문자열 연산 

#### 5.5.1. 문자열 추출 

- 슬라이싱을 사용하면 문자열을 추출 가능 (like substring)
- 1개만 추출하면 문자 번호를 추출

문자열 예제 1 : [string_oper1.go](section5/string_oper1.go)

#### 5.5.2. 문자열 비교 

- 같음(`==`), 같지 않음(`!=`) 비교는 가능
- 크기비교는 아스키코드값 기준으로 앞에서부터 비교함 (`<, >`)

문자열 예제 2 : [string_oper2.go](section5/string_oper2.go)

#### 5.5.3. 문자열 결합

- plus(`+`) 연산을 통해 단순 결합 가능 
- join 을 통하여 성능을 고려하여 결합 가능 

문자열 예제 3 : [string_oper3.go](section5/string_oper3.go)

## Section 6. 데이터타입 Advanced (array, slice, map, pointer)

### 6.1. 배열 (Array)

#### 6.1.1. 배열 선언 및 사용

- 배열은 용량, 길이 항상 같음 
- 배열 : 길이 고정, 값 타입, 복사 전달, 비교 연산자 사용 가능
- 슬라이스 : 길이 가변, 참조 타입, 참조 값 전달, 비교 연산자 사용 불가능 
- `cap()` : 배열, 슬라이스 용량 구하는 함수
- `len()` : 배열, 슬라이스 갯수 구하는 함수

배열 사용 예제 : [arra1.go](section6/array1.go)

#### 6.1.2. 배열 순회 

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

#### 6.1.3. 배열의 값 복사

- 배열을 복사하면 값 복사로 복사 됨

배열 값 복사 예제 : [array3.go](section6/array3.go)

### 6.2. 슬라이스 기초

#### 6.2.1. 슬라이스 선언 및 사용

- 슬라이스의 길이는 가변 
- 동적으로 크기가 늘어남 
- 레퍼런스(참조 값 타입)
- 선언 방법 :
  1. 배열처럼 선언 (길이는 명시하지 않고 사용)
  2. `make` 함수 (자료형, 길이, 용량(생략시 길이))

슬라이스 기초 예제 1 : [slice_basic1.go](section6/slice_basic1.go)

#### 6.2.2. 슬라이스 참조 복사 / 순회 

- 슬라이스를 복사하면 참조값이 복사 됨
- 메모리상의 동일한 슬라이스를 참조하는 것임
- 주의 : `make` 로 슬라이스 선언 시 길이 만큼만 초기화되기때문에 capa 를 별도로 지정한 경우는 panic 을 조심하자

슬라이스 기초 예제 2 : [slice_basic2.go](section6/slice_basic2.go)

### 6.3. 슬라이스 심화

#### 6.3.1. 슬라이스 추가 및 병합 

- 슬라이스에 요소나 다른 슬라이스를 추가하려면 `append` 사용
- 슬라이스의 capa 는 꽉 찼을경우 2배씩 증가 됨 

슬라이스 심화 예제 1 : [slice_ex1.go](section6/slice_ex1.go)

#### 6.3.2. 슬라이스 추출 및 정렬

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

#### 6.3.3. 슬라이스 값 복사 

- `copy` 로 복사 가능
- `make` 로 공간을 할당 후 복사 해야함. 공간이 할당되지 않으면 복사되지 않음
- 복사 된 슬라이스 의 값을 변경해도 원본에는 영향 없음 
- 주의 : 부분적인 슬라이스 추출은 참조 가 되어버림
- 부분 슬라이스 추출할때 3번째 값을 적어주면 참조되어지는 Target 슬라이스의 capa 를 지정할 수 있음 

슬라이스 심화 예제 3 : [slice_ex3.go](section6/slice_ex3.go)

### 6.4. 맵 (Map)

#### 6.4.1. 맵 - 선언과 사용

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

#### 6.4.2. 맵 - 순회 

- `range` 키워드를 활용한 `for` 문으로 순회 가능
- 맵은 순회할때 순서를 보장하지 않음

맵의 순회 예제 : [map2.go](section6/map2.go)

#### 6.4.3. 맵 - 값 추가 / 수정 / 삭제

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

#### 6.4.4. Map 조회시 키가 존재하는지 여부 확인 

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

### 6.5. 포인터 (Pointer)

#### 6.5.1. 포인터

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

## Section 7. 함수, Defer, Closure

### 7.1. 함수 Basic

#### 7.1.1. 함수 - 선언

- `func` 키워드로 선언
- 여러 개의 반환 값(return value) 사용 가능 

```go
func 함수명(매개변수) (반환타입 or 반환 값 변수명) : 반환 값 존재 
func 함수명() (반환타입 or 반환 값 변수명) : 매개변수 없음, 반환 값 존재 
func 함수명() : 매개변수 없음, 반환 값 없음
```

> [Golang 스펙 - Function types](https://go.dev/ref/spec#Function_types)

#### 7.1.2. 함수 - 선언 및 매개변수, 리턴값 사용

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

### 7.2. 함수 Advanced

#### 7.2.1. 함수 - 변수에 저장 (like function pointer)

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

#### 7.2.2. Defer 함수 (지연 호출 함수)

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


#### 7.2.3. Closure

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

## Section 8. 사용자 정의 타입, 구조체, 인터페이스 
- 상태, 메소드 분리해서 정의 (결합성 없음)
- 사용자 정의 타입 : 구조체, 인터페이스, 기본 타입(int, float, string...), 함수 
- 구조체와 메소드 연결을 통해서 타 언어의 클래스 형식처럼 사용 가능

### 8.1. type 키워드를 사용한 구조체선언, 기본타입 및 함수 재정의 

#### 8.1.1. 구조체의 선언 및 기본적인 사용

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

#### 8.1.2. 기본타입의 재정의 타입
- **무슨용도로 사용할까?**
```go
type cnt int

func myFunc() {
	var num cnt = 12
	...
}
```

#### 8.1.3. 함수의 재정의 타입
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

### 8.2. 구조체 (struct)

#### 8.2.1. 구조체 기초
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

#### 8.2.2. reflect 를 활용한 구조체 메타 정보 획득
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

#### 8.2.3. 중첩 구조체
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

### 8.3. 구조체 Advanced

#### 8.3.1. 생성자 패턴 사용

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

#### 8.3.2. 구조체의 메소드 호출 - 값참조, 주소참조

- 값 참조시에는 원래의 값 영향 받지 않음
- 주소 참조시에는 원래의 값 영향 받음 

구조체 메소드 호출 예제 : [struct_ex2.go](section8/struct_ex2.go)

#### 8.3.3. 메소드 오버라이딩
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

### 8.3. 인터페이스 (Interface)

#### 8.3.1. 인터페이스 기초 

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

#### 8.3.2. 덕 타이핑 (Duck typing)
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

#### 8.3.3. 빈 인터페이스 사용 
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

#### 8.3.4. Type Assertion
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
