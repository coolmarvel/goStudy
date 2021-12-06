package main

import "fmt"

const X int = 5

func main() {

	//배열을 선언시 갯수는 반드시 상수
	var num [5]int   //정수형 타입의 배열 5개
	fmt.Println(num) //디폴트값 0(초기화 안하면)
	var color [3]string
	fmt.Println(color)
	var color1 [3]string = [3]string{"red", "green", "blue"}
	color2 := [3]string{"a", "b", "c"}
	fmt.Println(color1)
	fmt.Println(color2)
	var fNum [5]float32 = [5]float32{1.5, 3.1} //0,1번째 인덱스는 각각 선언한 초기값 나머지는 0으로..
	fmt.Println(fNum)
	var num1 = [5]int{1: 20, 4: 100}
	fmt.Println(num1)
	//...를 이용해 배열의 갯수를 생략할수 있다.
	num2 := [...]int{10, 20, 30}
	fmt.Println(num2)

	// x := 5
	// a := [X]int{1, 2, 3, 4, 5}

	var nums = [...]int{1, 2, 3, 4, 5}
	fmt.Println(nums[4])

	//len:배열길이 리턴
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	var number [5]int = [5]int{1, 2, 3, 4, 5}
	for i, v := range number {
		fmt.Println(i, v) //i는 인덱스 ,v는 요소
	}
	fmt.Println("배열의 요소값만 출력하기")
	//인덱스가 필요없다면 _를 이용해 인덱스를 무효화
	for _, v := range number {
		fmt.Println(v) //요소만 출력
	}
	//배열은 연속된 메모리다.
	//컴퓨터는 인덱스와 타입크기를 사용해서 메모리 주소를 찾는다.

	//
	// a := [5]int{1, 2, 3, 4, 5}
	// b := [5]int{100, 200, 300, 400, 500}

	// for i, v := range a {
	// 	fmt.Printf("a[%d] = %d\n", i, v)
	// }
	// println()
	// for i, v := range b {
	// 	fmt.Printf("b[%d] = %d\n", i, v)
	// }
	// b = a //a배열을 b에 복사
	// println()
	// for i, v := range b {
	// 	fmt.Printf("b[%d] = %d\n", i, v)
	// }
	//다중배열

	b := [2][5]int{
		{1, 2, 3, 4, 5},  //b[0]
		{6, 7, 8, 9, 10}, //b[1]
	}

	for _, arr := range b { //arr순서대로 b[0],b[1]
		for _, v := range arr { //v값은 b[0],b[1]배열의 각 요소
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

}
