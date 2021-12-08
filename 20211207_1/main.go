package main

import "fmt"

type account struct {
	balance    int
	firstName  string
	secondName string
}

//포인터 메서드
func (a1 *account) PointerMethod(amount int) {
	a1.balance -= amount
}

//값 타입 메서드
func (a2 account) ValueMethod(amount int) {
	a2.balance -= amount
}

func (a3 account) ReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}
func main() {
	//리시버는 두가지 타입으로 정의할수 있다.
	//1.값타입 2. 포인터 타입

	var A *account = &account{200, "홍길동", "이길동"}
	A.PointerMethod(10)    //포인터 메서드 부르자
	fmt.Println(A.balance) //190

	A.ValueMethod(20)
	fmt.Println(A.balance) //190

	var B account = A.ReturnValue(30)
	fmt.Println(B.balance) //160

	B.PointerMethod(10)
	fmt.Println(B.balance) //150

	//리시버가 있으면 메서드, 없으면 걍 함수
	//리시버는 메서드를 호출하는 주체.메서드는 리시버를 통해서만 호출.
	//메서드는 리시버에 속한 기능이다.
	//포인터 메서드는 인스턴스 중심.값을 변경할수 있다.
	//값 타입은 값이 모두복사.복사하는 양에 따라서 성능이슈가 발생할수 있다.

}
