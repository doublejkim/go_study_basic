# Go Study - Basic

## Section 1

### 1.1. Hello World

[helloworld.go](section1/helloworld.go)

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
- 변수 선완과 마찬가지로 그룹화해서 선언 가능

여러개의 상수 선언 예제 : [const2.go](section2/const2.go)



