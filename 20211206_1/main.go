package main

import "fmt"

func main() {
	// slice : 동적 배열 -> 자동으로 배열 크기를 증가시켜준다.
	// 길이가 요소 갯수에 따라 자동으로 증가해 관리가 편함.
	// slicing 기능을 사용해 배열의 일부를 나타내는 slice를 만들 수 있다.

	// var arr[10]int -> 10개까지...

	// slice를 초기화하지 않으면 길이가 0인 slice가 만들어진다.
	var slice []int

	if len(slice) == 0 {
		fmt.Println("slice가 비어있다.", slice)
	}
	// slice[1] = 10 // 할당되지 않은 메모리 공간에 접근 -> 에러 발생
	// fmt.Println(slice)

	// {}를 사용한 초기화
	// var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3}
	fmt.Println(slice2)

	// var arr = [...]int{1, 2, 3} // 정적배열
	// var slice = []int{1, 2, 3} // slice(동적배열)

	// make()를 이용한 초기화
	// var slice3 = make([]int, 3) // 첫 번째 인수로 Type, 두 번째 인수로는 Len

	// slice 접근 : 배열이랑 같음.
	var slice4 = make([]int, 3)
	slice4[1] = 5 // [0, 5, 0]

	var slice5 = make([]int, 5)

	// slice 순회 -> 동적으로 길이가 늘어나는 거 말고 배열과 동일함.
	// 초기화
	for i := 0; i < len(slice5); i++ {
		slice5[i] = i + 1
	}
	// 출력
	for _, v := range slice5 {
		fmt.Println(v)
	}

	// 요소 추가
	var slice6 = []int{1, 2, 3}

	// 첫 번째 인수 : 추가하려는 slice, 두 번째 인수에는 추가하려는 요소
	slice7 := append(slice6, 4, 5, 7, 8) // 추가

	fmt.Println(slice6)
	fmt.Println(slice7)
	// append : 첫 번째 인수로 들어온 slice의 값을 변경하는게 아니라 요소가 추가된 새로운 slice를 반환
	// 기존 slice에 요소를 추가하고 싶다면 append() 결과를 기존 slice에 대입해서 변경해야 된다.

	var slice8 []int

	// for문으로 하나씩 추가
	for i := 0; i <= 10; i++ {
		slice8 = append(slice8, i)
	}
	slice8 = append(slice8, 11, 3, 123, 234)
	fmt.Println(slice8)

	// type SliceHeader struct {
	// 	Data uintptr // 실제 배열을 기리키는 포인터
	// 	Len  int     // 요소 갯수
	// 	Cap  int     // 실제 배열의 길이
	// }

}
