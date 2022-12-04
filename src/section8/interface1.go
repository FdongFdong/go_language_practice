package main

import "fmt"

type test interface { // 빈 인터페이스

}

func main() {
	// 인터페이스
	// 객체의 동작을 표현, 골격
	// 단순히 동작에 대한 방법만 표시
	// 추상화 제공
	// 인터페이스의 메서드를 구현한 타팁은 인터페이스로 사용 가능
	// Go 언어를 유연하게 사용 가능
	// 덕타이핑 : Go언어의 독창적인 특성
	// 인터페이스 -> 자바(전략패턴, 템플릿메서드, 팩토리메서드, 어댑터 패턴....)
	// 디자인 패턴 측면에서 Client 입장 -> 메서드의 구체적인 클래스를 몰라도 인터페이스 정의된 메서드를 사용하는 객체 보장
	// -> 클래스간의 결합도 감소 -> 유지보수성 향상, 기능 추가의 용이성 -> 독립적인 프로그래밍 가능

	// 예제 1
	/*
		type 인터페이스명 interface{
			메서드1() 반환 값(타입 형)
			메서드2() // 반환 값 없을 경우
		}
	*/

	var t test
	fmt.Println("ex1 : ", t) // 빈(Empty) 인터페이스인 경우 Nil 리턴

}