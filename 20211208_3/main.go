package main

import (
	"container/list"
	"fmt"
)

// type Element struct {
// 	Value interface{} //데이터를 저장
// 	Next  *Element    //다음 요소의 주소를 저장
// 	Prev  *Element    //이전 요소의 주소를 저장
// }

//큐 구조체
type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}
func (q *Queue) Pop() interface{} {
	front := q.v.Front()

	if front != nil {
		return q.v.Remove(front)
	}
	return nil

}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	//리스트 : 각데이터를 담고 있는 요소들을 포인터로 연결한 자료구조

	// v := list.New()       //리스트 생성
	// e4 := v.PushBack(4)   //뒤에 추가
	// e1 := v.PushFront(1)  //앞에 추가
	// v.InsertBefore(3, e4) //e4앞에 추가
	// v.InsertAfter(2, e1)  //e1다음에 추가

	// for e := v.Front(); e != nil; e = e.Next() { //요소 순회
	// 	fmt.Print(e.Value, "->")
	// }
	// fmt.Println()
	// for e := v.Back(); e != nil; e = e.Prev() { //역방향
	// 	fmt.Print(e.Value, "->")
	// }

	queue := NewQueue() //큐 만들기

	for i := 0; i < 5; i++ {
		queue.Push(i)
	}
	v := queue.Pop()
	for v != nil {

		fmt.Printf("%v->", v)
		v = queue.Pop()
	}
}

//지금 해볼것.리스트를 이용해서 스택만들기
