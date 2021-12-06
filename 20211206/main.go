package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func stringToUpper(str string) string {
	var a string
	for _, v := range str {
		if v >= 'a' && v <= 'z' {
			a += string('A' + (v - 'a')) // 합연산 이용
		} else {
			a += string(v)
		}
	}
	return a
}

// strings.Builder 객체를 이용해 문자를 더한다.
// strings.Builder는 내부에 slice를 가지고 있기 때문에
// WriteRune을 통해 문자를 더할 때 매번 메모리를 생성하지 않고 기존 메모리 공간에
// 빈자리가 있으면 그냥 더한다. 공간 낭비를 줄일 수 있다.
func stringToUpper1(str string) string {
	var b strings.Builder
	for _, v := range str {
		if v >= 'a' && v <= 'z' {
			b.WriteRune('A' + (v - 'a')) // strings.Builder 사용
		} else {
			b.WriteRune(v)
		}
	}
	return b.String()
}

func main() {
	// 문자열 합산
	var str string = "Hello"
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str)) // 내부 구조체로 변경
	address := stringHeader.Data                                  // Data 필드값을 변수로 저장
	str += "World"

	address1 := stringHeader.Data // Data 필드값을 변수로 저장
	str += "Hi"
	address2 := stringHeader.Data // Data 필드값을 변수로 저장

	fmt.Println(str)
	fmt.Printf("주소 : \t&x\n", address)
	fmt.Printf("주소1 : \t&x\n", address1)
	fmt.Printf("주소2 : \t&x\n", address2)
	// 위에서 알 수 있듯 기존 문자열 메모리 공간을 건드리지 않고 새로운 메모리 공간을 잡는다.
	// 이 말은 문자열 불변의 원칙을 지키게 된다.
	// 문자열 합연산이 빈번하게 발생되면 메모리가 낭비됨.
	// 만약 빈번히 사용된다면 string 패키지의 Builder를 이용해 메모리 낭비를 줄이도록 한다.

	var upperString string = "Hello world"
	fmt.Println(stringToUpper(upperString))
	fmt.Println((stringToUpper1(upperString)))
}
