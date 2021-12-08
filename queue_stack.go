package main

import "fmt"

type Queue []string

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(str string) {
	*q = append(*q, str)
}

func (q *Queue) Dequeue() (string, bool) {
	if q.IsEmpty() {
		return "", false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}

func main() {
	var queue Queue

	queue.Enqueue("구조자료")
	queue.Enqueue("And")
	queue.Enqueue("알고리즘")

	for !queue.IsEmpty() {
		x, y := queue.Dequeue()
		if y == true {
			fmt.Println(x)
		}
	}
}
