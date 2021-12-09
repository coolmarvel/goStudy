package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintChar() {
	char := []rune{'대', '한', '민', '국', '코', '리', '아'}

	for _, v := range char {
		time.Sleep(300 * time.Millisecond) //300ms
		fmt.Printf("%c", v)
	}
}
func PrintNum() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond) //400ms
		fmt.Printf("%d", i)
	}
}

var wg sync.WaitGroup

func Sum(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지의 합 %d.\n", a, b, sum)

	wg.Done() //작업이 완료됨
}

func main() {
	////////////////////////////////////////////////
	//go 함수이름->해당함수를 수행하는 새로운 고루틴 생성
	//호출된 함수는->새로운 고루틴에서 수행된다.

	go PrintChar() // 새로운 고루틴
	go PrintNum()

	time.Sleep(3 * time.Second) //3초간대기
	////////////////////////////////////////////////
	//sync패키지 WaitGroup

	// var wg sync.WaitGroup

	// wg.Add(5)  //작업 개수 설정
	// wg.Done() //작업이 완료 될때마다 호출
	// wg.Wait() //모든 작업이 완료 될때까지 대기

	wg.Add(10) //개수 설정

	for i := 0; i < 10; i++ {
		go Sum(1, 1000000)
	}
	wg.Wait() //모든 작업이 끝날때까지 기다리자
}
