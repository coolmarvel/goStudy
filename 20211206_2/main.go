package main

import "fmt"

func changeArr(arr [5]int) {
	arr[2] = 200
}

func changeSlice(slice []int) {
	slice[2] = 200
}

func main() {
	// var slice = make([]int, 3)

	array := [5]int{1, 2, 3, 4, 5}
	sli := []int{1, 2, 3, 4, 5}

	changeArr(array)
	changeSlice(sli)
	fmt.Println(array)
	fmt.Println(sli)

	slice1 := make([]int, 3, 5) // len : 3 cap : 5 slice 만듬
	slice2 := append(slice1, 4, 5)
	fmt.Println("slice1 :", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2 :", slice1, len(slice2), cap(slice2))

	/*
		만약 빈공간이 없으면 더 큰 배열을 만듬. 일반적으로 기존 배열의 2배
		=> 기존 배열의 요소를 새로운 배열에 복사 => 새로운 배열 맨뒤에 새 값을 추가
		cap : 새로운 배열의 길이 값
		len : 기존 길이에 추가한 개수만큼 더한 값
		pointer : 새로운 배열을 가리키는 slice 구조체를 return
	*/

	slice3 := []int{1, 2, 3}       // len : 3, cap : 3
	slice4 := append(slice3, 4, 5) // 길이가 6인 새로운 배열을 만들고 slice 배열의 모든 값을 복사 => 맨 뒤에 4, 5 len : 5, cap : 6

	// fmt.Println(("slice3 : ", slice3, len(slice3), cap(slice3)))
	// fmt.Println(("slice4 : ", slice4, len(slice4), cap(slice4)))

	slice3[1] = 200
	fmt.Println("slice3 : ", slice3, len(slice3), cap(slice3))

}

// // slicing : 배열의 일부를 집어낸다.
// 	// array[startIndex : endIndex]
// 	// 주의할 점 : endIndex -1 까지
// 	arr := [5]int{1,2,3,4,5}
// 	slice :=arr[1:2] // slicing

// 	fmt.Println("arr : ", arr)
// 	fmt.Println("slice : ", slice, len(slice), cap(slice))

// 	arr[1]=100 // arr의 두 번째 값 변경

// 	fmt.Println("arr : ", arr)
// 	fmt.Println("slice : ", slice, len(slice), cap(slice))

// 	slice=append(slice,200)

// 	fmt.Println("arr : ", arr)
// 	fmt.Println("slice : ", slice, len(slice), cap(slice))

// 	slice1:=[]int{1,2,3,4,5}
// 	slice2:=slice1[1:2] :=
