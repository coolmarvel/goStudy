package main

import (
	"fmt"
)

//인수 타입 앞에... 해당타입 인수를 여러개 받는 가변인수.
//가변 인수는 함수 내부에서 해당타입의 슬라이스로 처리
func sum(nums ...int) int { //가변 인수를 받는 함수
	sum := 0

	fmt.Printf("타입 : %T\n", nums)

	for _, v := range nums {
		sum += v
	}

	return sum

}

// func Print(args ...interface{}) string { //모든 타입을 받는 가변인수

// 	for _, arg := range args { //모든 인수를 돌면서
// 		switch b := arg.(type) { //인수 타입에 따라서
// 		case bool:
// 			val := arg.(bool) //변환
// 			//출력...
// 		case int:
// 			val := arg.(int) //변환
// 		}
// 	}
// }

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			//함수 리터럴을 이용해서 더하기 함수 정의하고 리턴
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			//함수 리터럴을 이용해서 곱하기 함수 정의하고 리턴
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	// fmt.Println()
	// fmt.Println(1)
	// fmt.Println(1,2,3,4,5,6)

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(sum(1, 2))
	fmt.Println(sum())

	fmt.Println(1, "ㅇㅇㅇ", 1.54142)

	////////////////////////////////////////////////////
	//defer 명령
	// f, err := os.Create("test.txt") //파일 생성
	// if err != nil {

	// 	fmt.Println("실패")
	// 	return
	// }

	// defer fmt.Println("호출")    //지연
	// defer f.Close()            //지연
	// defer fmt.Println("파일 닫음") //지연
	// fmt.Println("작성")
	// fmt.Fprintln(f, "Hello World")

	////////////////////////////////////////////////////

	//함수 리터럴 : 이름이 없는 함수->함수명을 쓰지 않고 함수 타입 변숫값으로 대입되는 함숫값, 익명함수나 람다
	//함수명이 없기 때문에 직접호출 X, 함수 타입변수로만 호출가능

	fn := getOperator("*")

	res := fn(3, 4) //함수 타입 변수를 사용해서 함수호출
	fmt.Println(res)

	////////////////////////////////////////////////////

	a := 0

	f := func() {
		a += 10
	}
	a++
	f() //함수 타입 변수
	fmt.Println(a)

}
