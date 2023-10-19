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