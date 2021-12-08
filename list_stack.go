package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()

	queue.PushBack("자료구조")
	queue.PushBack("And")
	queue.PushBack("알고리즘")

	front := queue.Front()
	fmt.Println(front.Value)
	queue.Remove(front)

	front = queue.Front()
	fmt.Println(front.Value)
	queue.Remove(front)

	front = queue.Front()
	fmt.Println(front.Value)
	queue.Remove(front)
}
