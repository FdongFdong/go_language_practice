# go_language_practice

- [go\_language\_practice](#go_language_practice)
  - [Go 언어 장점 및 특징](#go-언어-장점-및-특징)
  - [변수 및 상수](#변수-및-상수)
  - [제어문 및 반복문](#제어문-및-반복문)
    - [if](#if)
    - [switch](#switch)
    - [for](#for)
  - [패키지](#패키지)
  - [접근제어 및 Alias](#접근제어-및-alias)
  - [초기화 메드(init)](#초기화-메드init)
  - [데이터 타입](#데이터-타입)
    - [Bool](#bool)
    - [숫자형 기초](#숫자형-기초)
    - [숫자형 연산](#숫자형-연산)
    - [문자열 기초](#문자열-기초)
    - [문자열 연산](#문자열-연산)
  - [자료구조](#자료구조)
    - [배열](#배열)
    - [슬라이스](#슬라이스)
    - [맵(Map)](#맵map)
    - [포인터(Pointer)](#포인터pointer)
  - [함수](#함수)
    - [함수 기초](#함수-기초)
    - [함수 심화](#함수-심화)
    - [Defer](#defer)
    - [Closure](#closure)
  - [객체 지향](#객체-지향)
    - [사용자 정의 타입](#사용자-정의-타입)
    - [구조체 기초](#구조체-기초)
    - [구조체 심화](#구조체-심화)
    - [인터페이스 기초](#인터페이스-기초)
    - [인터페이스 심화](#인터페이스-심화)
  - [Go 병행 처리](#go-병행-처리)
    - [고루틴(Goroutine)](#고루틴goroutine)
    - [채널(Channel)](#채널channel)
  - [marshal unMarshal](#marshal-unmarshal)
  - [Ethereum Core Source를 이용한 Explorer 개발](#ethereum-core-source를-이용한-explorer-개발)

## Go 언어 장점 및 특징

- 간결한 문법 및 단순함
- 병행 프로그래밍 지원
- 정적 타입 및 동적 실행
- 간편한 협업 지원
- 컴파일 및 실행속도 빠름
- 제네릭 및 예외 처리 미지원
- 컨벤션 통일

___

## 변수 및 상수

> [variable1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/variable1.go)
>
> [variable2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/variable2.go)
>
> [variable3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/variable3.go)
>
> [const1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/const1.go)
>
> [const2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/const2.go)
>
> [enumeration1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/enumeration1.go)
>
> [enumeration2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/enumeration2.go)
>
> [enumeration3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/enumeration3.go)

- 기본 초기화
  - 정수 타입: 0,
  - 실수(소수점) : 0.0,
  - 문자열 : "",
  - Boolean: true, false
  - 변수명 : 숫자 첫글자x, 대소문자 구분o, 문자, 숫자 밑줄, 특수기호 사용 가능
  - 변수 및 상수 : 함수 내외 사용 가능
- 여러개 선언 하기
- 짧은 선언
  - 함수 안에서만 사용 가능(전역 x)
  - 선언 후 할당 시 예외 발생
  - 주로 제한된 범위의 함수 내에서 사용할 경우 코드 가독성을 높일 수 있다.
- 상수
  - const 예약어를 통해 사용 초기화
  - 한 번 선언 후 값 변경 금지
  - 고정 된 값 관리용
- 열거형
  - 상수를 사용하는 일정한 규칙에 따라 숫자를 계산하거나 증가시키는 묶음

## 제어문 및 반복문

### if

> [if1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/if1.go)
>
> [if2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/if2.go)
>
> [if3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/if3.go)

- 괄호는 if문 뒤에 둬야한다.

### switch

> [switch1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/switch1.go)
>
> [switch2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/switch2.go)
>
> [switch3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/switch3.go)

- Switch 키위드 뒤 표현식(Expression) 생략 가능
- case 뒤 표현식(Expression) 사용 가능
- 자동 break 때문에 fallthrough 존재
- Type 분기 -> 값이 아닌 변수 Type으로 분기 가능

### for
> [for1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/for1.go)
>
> [for2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/for2.go)
>
> [for3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/for3.go)

- while, do while문이 없다.
- Go에서 유일하게 제공되는 반복문
- 다양한 사용법 숙지 필요


___

## 패키지

> [package1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/package1.go)
>
> [package2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/package2.go)

- 패키지 : 코드 구조화(모듈화) 및 재사용
- 응집도, 결합도
- Go : 패키지 단위의 독립적이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그램을 작성할 것을 권고
- 패키지 이름 = 디렉토리 이름
- 같은 패키지 내 -> 소스 파일들은 디렉토리명을 패키지 명으로 사용한다.
- 네이밍 규칙 : 소문자 private, 대문자 : public
- Go : main 패키지는 특별하게 인식 -> 컴파일러 공유 라이브러리 x, 프로그램의 시작점 start point

___

## 접근제어 및 Alias

> [access1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/access1.go)
> 
> [access2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/access2.go)

- 패키지 접근제어
- 변수, 상수, 함수, 메서드, 구조체 등 식별자
- 첫글자가 대문자 : 패키지 외부에서 접근가능
- 첫글자가 소문자 : 패키지 외부 접근 불가(패키지 내에서만 접근가능)
- 별칭 사용
- 빈 식별자 사용

___

## 초기화 메드(init)

> [init1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/init1.go)
>
> [init2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/init2.go)
>
> [init3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/init3.go)

- init : 패키지 로드 시에 가장 먼저 호출된다.
- 가장 먼저 초기화 되는 작업 작성 시 유용하다.
- init()는 여러개 있어도 컴파일된다. 
- 다른 패키지에 있으면 해당 패키지에 있는 init함수가 가장 먼지 실행된다.
![image](https://user-images.githubusercontent.com/20445415/204508694-08e23dcc-76d6-41c1-86d0-0cb1147e8b07.png)

___

## 데이터 타입

### Bool

> [bool1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/bool1.go)
> 
> [bool2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/bool2.go)

- Bool(Boolean) : 참과 거짓
- 조건부 논리 연산자랑 주로 사용 : !, ||(or), &&(and)
- 암묵적 형변환이 안된다. : 0, Nil -> False 변환 없음

### 숫자형 기초

> [numeric1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/numeric1.go)
>
> [numeric2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/numeric2.go)
>
> [numeric3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/numeric3.go)

- 데이터 타입 : 숫자형
- 정수, 실수, 복소수
- 32bit, 64bit, unsigned(양수)
- 정수 : 8진수(0), 16진수(0x), 10진수

### 숫자형 연산

> [number_operation1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/number_operation1.go)
>
> [number_operation2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/number_operation2.go)
>
> [number_operation3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/number_operation3.go)

- 숫자 연산(산술, 비교)
- 타입이 같아야 가능하다.
- 다른 타입끼리는 반드시 형 변환 후 연산
- 형 변환이 없을 경우 예외(에러) 발생
- +, -, *, %, /, <<, >>, &, ^

### 문자열 기초

> [string1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/string1.go)
>
> [string2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/string2.go)

- 큰 따옴표로 감싼다 "" 또는 백스쿼트 ``
- Golang : 문자 char 타입 존재하지 않음 -> rune(int32)로 문자 코드 값으로 표현
- 문자 : ''로 작성 가능
- 자주 사용하는 escape : \\, \', \", \a(콘솔벨), \b(백스페이스), \f(쪽바꿈), \n(줄바꿈), \r(복귀), \t(탭)...
- 문자열 : UTF-8 인코딩(유니코드 문자 집합) -> 바이트 수 주의

### 문자열 연산

> [string_oper1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/string_oper1.go)
>
> [string_oper2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/string_oper2.go)
>
> [string_oper3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/string_oper3.go)

- 문자열 연산
- 추출, 비교, 조합(결합)

___

## 자료구조

### 배열

> [array1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/array1.go)
>
> [array2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/array2.go)
>
> [array3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/array3.go)

- 배열은 용량, 길이가 항상 같다.
- 배열 vs 슬라이스 차이점 중요
- 길이 고정 vs 길이 가변
- 값 타입 vs 참조 타입
- 복사 전달 vs 참조 값 전달
- 전체 비교연산자 사용가능 vs 비교 연산자 사용불가
- 대부분 슬라이스 사용한다.
- cap() : 배열, 슬라이스 용량 계산
- len() : 배열, 슬라이스 개수 계산

### 슬라이스

> [slice_basic1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/slice_basic1.go)
>
> [slice_basic2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/slice_basic2.go)
>
> [slice_ex1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/slice_ex1.go)
>
> [slice_ex2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/slice_ex2.go)
>
> [slice_ex3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/slice_ex3.go)

- 길이 가변 -> 동적으로 크기가 늘어난다.
- 레퍼런스 타입(참조 값)
- 길이와 용량의 개념이 있다
- 크키가 동적으로 할당된다.
- 2가지 선언 방법
  - 배열처럼 선언
  - make 함수 선언 : make(자료형, 길이, 용량(생략시 길이값이 대입된다.))
- 슬라이스 추출 및 정렬
- slice[i:j] : i -> j-1 까지 추출
- slice[i:] : i -> 마지막까지 추출
- slice[:j] : 처음 -> j-1 까지 추출
- slice[:] : 처음 -> 마지막까지 추출
- 슬라이스 복사
  - copy(복사 대상, 원본)
  - 반드시 make로 공간을 할당 후 복사해야 한다.
  - 복사 된 슬라이스 값 변경해도 원본에는 영향 없음

### 맵(Map)

> [map1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/map1.go)
>
> [map2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/map2.go)
>
> [map3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/map3.go)
>
> [map4.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/map4.go)

- 맵 : 해시테이블, 딕셔너리(파이썬)
- key - value로 자료 저장
- 레퍼런스 타입(참조 값 전달)
- 비교 연산자 사용 불가능(참조 타입이므로)
- 특징 : 참조타입(key)로 사용 불가능, 값(value)으로 모든 타입 사용가능
- make 함수 및 축약(리터럴)로 초기화 가능
- 순서 없음
- 맵 조회 및 순회(Iterator)
  - 순서가 없으므로 랜덤

### 포인터(Pointer)

> [pointer1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/pointer1.go)
>
> [pointer2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/pointer2.go)
>
> [pointer3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section6/pointer3.go)

- Go : 포인터 지원(C)
- 변수의 지역성, 연속된 메모리의 참조성 ..., 힙, 스택..
- 주소의 값은 직접 변경 불가능(잘못된 코딩으로 인한 버그 방지)
- *(애스터리스크)로 사용
- nil로 초기화(nil != 0)
- 포인터로 값 전달 시
  - 함수, 메서드 호출 시 매개변수 값을 복사해서 전달한다. -> 함수, 메서드 내에서는 원본 값 변경이 불가능하다.
  - 원본 값 변경 위해서는 포인터로 전달
  - 특히 크기가 큰 배열인 경우 값 복사시 시스템에 부담이간다 -> 포인터 전달로 해결(슬라이스, 맵은 참조 전달)

___

## 함수

### 함수 기초

> [func1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/func1.go)
>
> [func2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/func2.go)
>
> [func3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/func3.go)
>
> [func4.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/func4.go)

- 선언 : func 키워드로 선언
  - func 함수명(매개변수) (반환타입 or 반환 값 변수명) : 반환 값 존재
  - func 함수명() (반환타입 or 반환 값 변수명) : 매개변수 없음, 반환 값 존재
  - func 함수명(매개변수) : 매개변수 존재, 반환 값 없음
  - func 함수명() : 매개변수 없음, 반환 값 없음
- 타 언어와 달리 반환 값(return value)이 다수 가능

### 함수 심화

> [func_ex1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/func_ex1.go)
>
> [func_ex2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/func_ex2.go)
>
> [func_ex3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/func_ex3.go)
>
> [func_ex4.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/func_ex4.go)

- 가변 인자 실습(매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않음)
- 함수를 변수에 할당할 수 있다.
- 익명함수를 사용할 수 있다.

### Defer

> [f_defer1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/f_defer1.go)
>
> [f_defer2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/f_defer2.go)
>
> [f_defer3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/f_defer3.go)
>
> [f_defer4.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/f_defer4.go)

- Defer 함수 실행(지연)
- Defer를 호출한 함수가 종료되기 직전에 호출된다.
- 타 언어의 Finally 문과 비슷
- 주로 리소스 반환, 열린 파일 닫기, Mutex 잠금 해제

### Closure

> [closure1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/closure1.go)
>
> [closure2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section7/closure2.go)

- 익명함수를 사용할 경우 함수를 변수에 할당해서 사용가능
- 함수 안에서 함수 선언 및 정의가능 -> 이 때 외부 함수에 선언된 변수에 접근 가능한 함수
- 함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근학기 위한 스냅샷(객체)
- 함수를 호출 할 때 이전에 존재했던 값을 유지하기 위해서 -> 비동기, 누적카운트 -> 무분별한 전역변수 남발..
- 무분별한 전역변수 남발 -> 객체들이 메모리에 존재하므로 -> 메모리부족, 오버플로우 현상 등 자원을 무분별하게 사용 가능성
- 클로저를 정확하게 이해하고 사용할 필요가 있음

___

## 객체 지향

- Go -> 객체 지향 타입을 구조체로 정의한다. (클래스, 상속 개념 없음)
- 객체지향 -> 클래서(속성 : 멤버변수, 기능(상태 : 메서드)) : 코드의 재사용성, 코드의 관리, 신뢰성이 높은 프로그래밍
- Go는 전형적인 객제지향의 전형적인 특징을 가지고 있지 않지만, 인터페이스 -> 다형성 지원, 구조체를 클래스형태의 코딩 가능
- 객체지향의 기본 개념 -> Go에서 포함하고 있다. -> 객체 지향 프로그래밍 언어
- 상태, 메서드 분리해서 정의(결합성 없음)
- 사용자 정의 타입: 구조체, 인터페이스, 기본 타입(int, float, string...), 함수
- 구조체와 메서드 연결을 통해서 타 언어의 클래스 형식처럼 사용가능(객체 지향)

### 사용자 정의 타입

> [user_struct1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/user_struct1.go)
>
> [user_struct2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/user_struct2.go)
>
> [user_struct3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/user_struct3.go)
>
> [user_struct4.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/user_struct4.go)

### 구조체 기초

> [struct1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/struct1.go)
>
> [struct2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/struct2.go)
>
> [struct3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/struct3.go)
>
> [struct4.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/struct4.go)
>
> [struct5.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/struct5.go)
>
> [struct6.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/struct6.go)

- Go의 필드들의 집합체 또는 컨테이너
- 필드는 갖지만 메서드는 갖지 않는다. 
- 객체지향 방식을 지원 -> 리시버를 통해 메서드랑 연결
- 상속, 객체, 클래스 개념 없음
- 구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)
- 다양한 선언 방법
  - &struct, struct : &struct 포인터를 받아오고 역참조를 또 하기 떄문에 속도가 조금 느리다
  - 인터페이스 메서드를 선언만 해둔 후 -> 오버라이딩 해서 메서드에 포인터 리시버를 사용할 경우에는 반드시 &struct

### 구조체 심화

> [struct_ex1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/struct_ex1.go)
>
> [struct_ex2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/struct_ex2.go)
>
> [struct_ex3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/struct_ex3.go)
>
> [struct_ex4.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/struct_ex4.go)
>
> [struct_ex5.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/struct_ex5.go)

- 구조체 임베디드 패턴
- 다른 관점으로 메서드를 재 사용하는 장점 제공
- 상속을 허용하지 않는 Go 언어에서 메서드 상속 활용을 위한 패턴
- 구조체 임베디드 메서드 오버라이딩 패턴

### 인터페이스 기초

> [interface1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/interface1.go)
>
> [interface2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/interface2.go)
>
> [interface3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/interface3.go)
>
> [interface4.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/interface4.go)
>
> [interface5.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/interface5.go)

- 객체의 동작을 표현, 골격
- 단순히 동작에 대한 방법만 표시
- 추상화 제공
- 인터페이스의 메서드를 구현한 타팁은 인터페이스로 사용 가능
- Go 언어를 유연하게 사용 가능
- 덕타이핑 : Go언어의 독창적인 특성
- 인터페이스 -> 자바(전략패턴, 템플릿메서드, 팩토리메서드, 어댑터 패턴....)
- 디자인 패턴 측면에서 Client 입장 -> 메서드의 구체적인 클래스를 몰라도 인터페이스 정의된 메서드를 사용하는 객체 보장
- -> 클래스간의 결합도 감소 -> 유지보수성 향상, 기능 추가의 용이성 -> 독립적인 프로그래밍 가능

### 인터페이스 심화

> [interface_ex1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/interface_ex1.go)
>
> [interface_ex2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/interface_ex2.go)
>
> [interface_ex3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section8/interface_ex3.go)

- 인터페이스 활용(빈 인터페이스)
  - 빈 인터페이스 : 함수 매개변수, 리턴 값, 구조체 필드등으로 사용 가능 -> 어떤 타입으로도 변환 가능
  - 모든 타입을 나타내기 위해 빈 인터페이스를 사용한다.
  - 동적타입으로 생각하면 쉽다. (타 언어의 Object 타입)
- Type Assertion
  - 실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
  - 인터페이스.(타입) -> 형 변환
  - interfaceVal.(type)
- 실제 타입 검사 switch 사용
  - 빈 인터페이스는 어떠한 자료형도 전달 받을 수 있으므로, 타입 체크를 통해 형 변환 후 사용 가능

___

## Go 병행 처리

### 고루틴(Goroutine)

> [goroutine1.go]()
>
> [goroutine2.go]()
>
> [goroutine3.go]()
>
> [goroutine4.go]()

- 타 언어의 쓰레드(Thread)와 비슷한 기능
- 생성 방법 매우 간단하며 리소스를 매우 적게 사용한다.
  - -> 수 많은 고루틴 동시 생성 실행 가능
- 비동기적 함수 루틴 실행(매우 적은 용량 차지) -> 채널을 통한 통신 가능
- 공유메모리 사용 시에 정확한 동기화 코딩 필요
- 싱글루틴에 비해 항상 빠른 처리 결과는 아니다.
- 멀티 쓰레드 장점과 단점
  - 장점
    - 응답성 향상
    - 자원 공유를 효율적으로 활용
    - 작업이 분리되어 코드 간결해짐
  - 단점
    - 구현이 어려움
    - 테스트 및 디버깅이 어려움
    - 전체 프로세스의 사이드 이펙트가 나올 수 있다.
    - 성능 저하
    - 동기화 코딩 반드시 숙지하여 사용해야한다
    - 데드락

### 채널(Channel)

> [gochannel1.go]()
>
> [gochannel2.go]()
>
> [gochannel3.go]()
>
> [gochannel4.go]()
>
> [gochannel5.go]()
>
> [gochannel6.go]()

- 고루틴간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해 사용 : 채널(동기식, 레퍼런스 타입)
- 실행 흐름 제어 가능(동기, 비동기) -> 일반 변수로 선언 후 사용 가능
- 데이터 전달 자료형 선언 후 사용(지정된 타입만 주고 받을 수 있음)
- interface{} 전달을 통해 자료형 상관없이 전송 및 수신 가능
- 값을 전달(복사 후 : bool, int 등), 포인터(슬라이스, 맵)등을 전달시에는 주의 > 동기화 사용(Mutex)
- 멀티 프로세싱 처리에서 교착상태(경쟁) 주의
- <- , -> ( 채널 <- 데이터 ) : 송신, ( <- 채널 ) : 수신

## marshal unMarshal

> [json](https://github.com/FDongFDong/go_language_practice/blob/main/json/json.go)
___

## Ethereum Core Source를 이용한 Explorer 개발

- Block event 데이터 처리하기
[go-ethereum/ethclient](https://github.com/ethereum/go-ethereum/tree/master/ethclient)
[go-ethereum/core/types](https://github.com/ethereum/go-ethereum/tree/master/core/types)

> [BlockListner.go](https://github.com/FDongFDong/go_language_practice/tree/main/explorer)

- Block 정보 가져오기
- Transaction 정보 가져오기
- ERC20 토큰 정보 가져오기

<img width="1765" alt="image" src="https://user-images.githubusercontent.com/20445415/201647346-b1a44e7d-a701-49a4-ba6a-e1b2dff84a40.png">
