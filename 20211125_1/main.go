package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Age   int32
	Score float64
}

type Test struct {
	A int8 //1
	B int  //8
	C int8 //1
	D int  //8
	E int8 //1

}
type Test1 struct {
	A int8 //1
	C int8 //1
	E int8 //1
	B int  //8
	D int  //8

}

func PrintStudent(st Student) {
	fmt.Printf("나이 :%d,점수 : %.2f", st.Age, st.Score)
}

func main() {

	var s = Student{15, 20.5}
	//PrintStudent(s)
	//s1 := s          //s 구조체 모든 멤버(필드)들이 s1복사
	//PrintStudent(s1) //함수 호출할때도 구조체가 복사

	fmt.Println(unsafe.Sizeof(s))
	var t = Test{1, 2, 3, 4, 5}
	var t1 = Test1{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(t))
	fmt.Println(unsafe.Sizeof(t1))

}
