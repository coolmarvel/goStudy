/*

패키지선언 :  이코드가 어떤 패키지에 속하냐?
go언어의 모든 코드는 반드시 패키지 선언으로 시작
package main : main패키지에 속한 코드다!!라는걸 컴파일러한테 알려줌
go언어는 패키지로 시작하고 package main은 프로그램 시작점이 있는 패키지다!!!

*/
package main

import (
	"fmt"
	"unsafe"
)

/*
	클래스X 메서드를 가지는 구조체를 지원
	상속 X
	메서드 O
	인터페이스 O
	익명함수O
	가비지 컬렉터O
	포인터
	제네릭프로그래밍 X
	네임스페이스 X 패키지 단위로 구분
*/

func main() {
	fmt.Print("나는 프린트 함수")
	fmt.Println("Hello World")
	fmt.Print("나는 프린트 함수")
	fmt.Printf("나는 프린트 함수\n")
	//var : 변수를 선언하는 키워드
	// int : 정수형 타입
	var a int = 10

	var num = 5 //우변의 있는 값을 추론
	fmt.Println("변수 num 의 사이즈", unsafe.Sizeof(num))

	num2 := 10 //선언대입문 :  :=를 써서 var키워드,타입을 생략
	fmt.Println("var키워드와 타입생략 방법", num2)

	//32비트 환경에서 int는 int32
	//64비트 환경에서 int는 int64

	//var b float32 = 20.005

	//변수는 항상 동일하게
	//var plus int = a + b->타입이 다르므로 안됨
	var msg string = "Hello"
	fmt.Println(a)
	fmt.Println(msg)
	//fmt.Println(plus)

	//변수속성은 네가지가 이씀. 이름, 값, 타입, 주소

	//숫자타입
	/*
		uint8 : 1바이트, 부호가 없는 정수
		uint16
		uint32
		uint64

		//실수
		float32
		float64

		//복소수
		complex64
		complex128


		byte uint8의 별칭
		rune int32의 별칭 UTF-8문자 하나를 나타낼때 사용
		bool
		string

	*/
	//0000	0000

	//Sizeof
	var int_size int
	var int_size8 uint8
	var int_size16 int16
	var int_size32 int32
	var int_size64 int64
	var uint_size64 uint64

	fmt.Println("int형 사이즈 : ", unsafe.Sizeof(int_size), "바이트")
	fmt.Println("int8형 사이즈 : ", unsafe.Sizeof(int_size8), "바이트")
	fmt.Println("int16형 사이즈 : ", unsafe.Sizeof(int_size16), "바이트")
	fmt.Println("int32형 사이즈 : ", unsafe.Sizeof(int_size32), "바이트")
	fmt.Println("int64형 사이즈 : ", unsafe.Sizeof(int_size64), "바이트")
	fmt.Println("uint64형 사이즈 : ", unsafe.Sizeof(uint_size64), "바이트")

	//타입변환(형변환)

	// num3 := 3
	// var fnum float64 = 3.5

	// var c int = int(fnum) //타입이 틀리기 때문에 대입도 되지 않는다.

	// res := num3 * fnum //타입이 다르기 때문에 연산이 되지 않는다.

	// var e int64 = 7

	// res1 := num3 * e //둘다 정수형이지만 타입이 틀리기 때문에 연산이 되지 않는다.

	/*
		00000000 00000000 00000001 01101000	->360
								   01101000  ->104
	*/
	var BigType int32 = 360
	var smallType int8 = int8(BigType)

	fmt.Println(smallType)

	{
		var s int = 50
		fmt.Println(s)
	}
	//fmt.Println(s)

	var defaultValueInteager int8
	var defaultValuefloat float32
	var isCheck bool
	var str string

	//초기값을 넣지 않으면 아래와 같이 디폴트값이 들어감
	//그외...nil
	fmt.Println("초기값을 넣지 않으면 ?디폴트값이 슉슉~ : ", defaultValueInteager)
	fmt.Println("초기값을 넣지 않으면 ?디폴트값이 슉슉~ : ", defaultValuefloat)
	fmt.Println("초기값을 넣지 않으면 ?디폴트값이 슉슉~ : ", isCheck)
	fmt.Println("초기값을 넣지 않으면 ?디폴트값이 슉슉~ : ", str)

}
