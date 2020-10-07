package main

import (
	"fmt"
)

type queueNode struct {
	Next  *queueNode
	Value interface{}
}

// Queue is Queue Type
type Queue struct {
	Head *queueNode
	Size int
}

func newQueue() *Queue {
	return &Queue{Head: &queueNode{Value: nil, Next: nil}, Size: 0}
}

func (q *Queue) push(v interface{}) {
	// 如果是第一次添加元素
	if q.Size == 0 {
		qn := &queueNode{Value: v, Next: nil}
		q.Head.Next = qn
		q.Size++
		return
	}
	node := q.Head
	for node.Next != nil {
		node = node.Next
	}
	qn := &queueNode{Value: v, Next: nil}
	node.Next = qn
	q.Size++
}

func (q *Queue) pop() (node *queueNode) {
	if q.Size == 0 {
		return nil
	}
	node = q.Head.Next
	q.Head.Next = node.Next
	q.Size--
	return
}

func main() {

	queue := newQueue()
	queue.push("1")
	queue.push("2")
	queue.push("3")
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())

}
