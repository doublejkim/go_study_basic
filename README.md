# Go Study - Basic

## Section 1

### 1.1. Hello World

[helloworld.go](section1/helloworld.go)

---

## Section 2

### 2.1. 변수  

#### 2.1.1. 변수 선언 및 사용

- `var` 로 선언 및 사용
- 값을 초기화 하지않으면 데이터 타입을 같이 정의 해야함 
- 값을 초기화 하면서 선언할 경우는 초기화 값을 보고 데이터타입을 추론하기때문에 사용 가능 

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

---

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
