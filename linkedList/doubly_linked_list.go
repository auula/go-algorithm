package main

import "fmt"

// 双向链表实现
// doubly linked list

type node struct {
	per   *node
	value interface{}
	next  *node
}

type doublyList struct {
	head   *node
	tail   *node
	length uint
}

func (n *node) Before() *node {
	return n.per
}

func (n *node) After() *node {
	return n.next
}

func (n *node) Element() interface{} {
	return n.value
}

func NewNode(v interface{}) *node {
	return &node{value: v}
}

func createDoubly() *doublyList {
	return &doublyList{}
}

func (d *doublyList) RPush(v interface{}) {
	// 链表未空的时候 创建链表
	head := d.head
	if head == nil {
		newNode := NewNode("this is hard node. ROOT node")
		d.head = newNode
		d.tail = newNode
		d.length++
	}
	for head != nil {
		head = head.next
	}
	newNode := NewNode(v)
	newNode.per = d.tail
	head.next = newNode
	d.tail = newNode
	d.length++
}

func (d *doublyList) getNode(index int) *node {
	node := d.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func main() {
	doubly := createDoubly()
	doubly.RPush("C++")
	doubly.RPush("Java")
	doubly.RPush("Golang")
	fmt.Println(doubly.length)
	fmt.Println("头节点是:", doubly.head)
	fmt.Println("尾节点是:", doubly.tail)
	fmt.Println(doubly.getNode(3))
}
