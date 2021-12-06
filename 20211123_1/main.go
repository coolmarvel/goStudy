package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}
func PrintScore(name string, math int, eng int, science int) {
	total := math + eng + science
	avg := total / 3
	fmt.Println(name, "의 평균 점수", avg, "이다")
}

//멀티반환함수(여러개 리턴값을 가지고 있음)
//매개변수 타입이 같으면 아래와 생략가능
// func Divide(a, b int) (int, bool) {
// 	if b == 0 {
// 		return 0, false
// 	}
// 	return a / b, true
// }

//변수명을 이용한 반환
func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return //명시적으로 반환할 값을 지정하지 않음
	}
	result = a / b
	success = true
	return

}

//피보나치 수열 :  어떤 수열의 항이 앞의 두항의 합과 같은 수열임
//ex)1,1,2,3,5,8,13,21,34,55,89,144,233...
//재귀
func Fibo(n int) int {

	//재귀 함수의 탈출 조건->없으면 무한 반복....반드시 있어야 함. 없으면 프로그램 뻗어버림.
	if n < 2 {
		return n
	}
	//나는 재귀->함수안에서 자기 자신 함수를 호출함.
	return Fibo(n-2) + Fibo(n-1)
}

func main() {
	c := Add(3, 6) //함수 호출
	fmt.Println(c)

	PrintScore("홍길동", 80, 80, 100)
	PrintScore("배추도사", 100, 80, 100)
	PrintScore("은비까비", 80, 70, 100)

	a, success := Divide(9, 3)
	fmt.Println(a, success)

	//지금 해볼것.
	//함수를 이용해서 평균값을 구해서
	//학점 뽑기
	//90~100 A
	//80~89 B
	//70~79 C
	//60~69 D
	//60점 미만은  걍 낙제

	//피보나치 수열의 9번째 값 : 34
	fmt.Println(Fibo(9))

}
