package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// type StringHeader struct{
// 	Data uintptr	//문자열의 데이터가 있는 메모리 주소를 나타내는 일종의 포인터
// 	Len int		//int타입의 문자열의 길이
// }
func main() {
	str := "안녕 난 홍길동이야"
	str1 := str
	fmt.Println(str)
	fmt.Println()
	fmt.Println(str1)
	//str의 Data와 Len값만 str1에 복사한다.

	// StringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str))
	// StringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str1))

	// fmt.Println(StringHeader1)
	// fmt.Println(StringHeader2)

	var str2 string = "Hello World"
	//str2 = "경일게임아카데미"
	// fmt.Println(str2)
	// str2[3] = 'v'
	//[]byte//-> 부호가 없는 정수타입의 가변길이 배열(1byte)

	var slice []byte = []byte(str2)

	slice[2] = 't'
	fmt.Println(str2)
	fmt.Printf("%s\n", slice)
	StringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
	StringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str2 : \t%x\n", StringHeader1.Data)
	fmt.Printf("slice : \t%x\n", StringHeader2.Data)

}
