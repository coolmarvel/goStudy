package main

import "fmt"

type account struct {
	b int
}

//일반함수
func Function(a *account, amount int) {
	a.b -= amount
}

//메서드방식
func (a *account) Method(amount int) {
	a.b -= amount

}

//사용자 정의
type myNum int

func (a myNum) add(b int) int {
	return int(a) + b
}

func main() {
	//method : 함수의 일종임. 고에서는 class X
	//구조체 밖에 메서드 지정.
	//어느 구조체에 속하냐?표시를 해줘야 함.
	//리시버 사용하면 됨.->메서드가 속하는 타입을 알려주는 녀석.

	//메서드 선언하기
	// func(r Rect) cal()int{
	// 	return r.width* r.height
	// }

	test := &account{200} //b가200
	Function(test, 30)

	test.Method(20)

	fmt.Printf("%d \n", test.b)

	var a myNum = 10
	fmt.Println(a.add(20))
	var b int = 20
	fmt.Println(myNum(b).add(100))

}
