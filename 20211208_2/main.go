package main

import "fmt"

func cap1() {
	f := make([]func(), 3) //함수 리터럴 3개를 가진 슬라이스
	fmt.Println("캡1")
	for i := 0; i < 3; i++ {

		f[i] = func() {
			fmt.Println(i) //
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}
func cap2() {
	f := make([]func(), 3) //함수 리터럴 3개를 가진 슬라이스
	fmt.Println("캡2")
	for i := 0; i < 3; i++ {
		v := i //v변수에 i값 복사

		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}

}
func main() {
	//함수 리터럴 외부 변수를 내부로 가져오는거를 캡쳐
	//캡쳐는 값복사가 아니라 레퍼런스(참조) 형태로 가져옴.

	cap1()
	cap2()
}
