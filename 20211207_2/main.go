package main

import "fmt"

type Test interface {
	String() string
}
type Student struct {
	Name string
	Age  int
}

func (s Student) String() string { //Student의 String()메서드

	return fmt.Sprintf("하이 %d, %s", s.Age, s.Name)
}

//빈 인터페이스
//빈 인터페이스를 인수를 받는 함수
func PrintValue(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf(" int %d\n", int(t))
	case float64:
		fmt.Printf(" float64 %f\n", float64(t))
	case string:
		fmt.Printf(" string %s\n", string(t))
	default:
		fmt.Printf(" 위에 없는 다른 타입: %T:%v\n", t, t)

	}
}

type People struct {
	Age int
}
type Mover interface {
	Moving()
}

func main() {
	//interface : 구현을 포함하지 않는 메서드의 집합
	//구체화된 타입 아닌 인터페이스만 가지고 메서드를 호출->프로그램 요구사항 변경시 유연하게 대처가능
	//Go->인터페이스 구현 여부를 그 타입이 인터페이스에 해당하는 메서드를 가지고 있는지로
	//판단하는 덕 타이핑을 지원한다.
	//타입 ->인터페이스 이름 써주고 ->인터페이스 키워드

	//인터페이스도 타입중 하나임.
	//인터페이스 변수 선언 가능. 변수의 값으로도 쓸수 있다.
	// type DuckInterface interface {
	// 	//메서드의 집합
	// 	//주의 할점: 메서드는 반드시 메서드명이 있어야함.
	// 	//매개변수와 리턴이 다르더라도 이름이 같은 메서드는 X
	// 	//인터페이스에서는 메서드 구현을 포함하지 않는다.
	// 	walk()
	// 	Move(distance int) int
	// }

	// type Test interface {
	// 	String() string
	// 	//String(int) string
	// 	//_(a int)
	// }

	st := Student{"이순신", 10} //Student타입
	var t Test               //Test타입

	t = st //t 값으로 st대입

	fmt.Printf("%s\n", t.String())

	//덕 타이핑
	//타입선언->인터페이스 구현여부를 명시적으로 나타낼 필요없이 인터페이스에
	//정의한 메서드 포함 여부만으로 결정하는 방식

	//인터페이스의 또다른 기능
	//1.인터페이스 기본값
	//2.빈 인터페이스
	//3.포함된 인터페이스

	// type Reader interface {
	// 	Read() (n int, err error)
	// 	Close() error
	// }
	// type Writer interface{
	// 	Write()(n int, err error)
	// 	Close() error
	// }
	// type ReadWriter interface{
	// 	Reader
	// 	Writer
	// }

	//빈 인터페이스
	//메서드를 들고 있지 않은 빈 녀석
	//이녀석은 어떤값이든 받을수 있는 함수,메서드,변수값을 만들때 사용
	PrintValue(10)
	PrintValue(3.141592)
	PrintValue("빈 인터풰이스")
	PrintValue(People{100})

	var m Mover
	m.Moving()

}
