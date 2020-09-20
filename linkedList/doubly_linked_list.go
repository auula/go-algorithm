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

func createDoublyList() *doublyList {
	d := new(doublyList)
	n := new(node)
	d.head = n
	d.tail = n
	d.length = 0
	return d
}

func (d *doublyList) RPush(v interface{}) {
	// 链表未空的时候 创建链表
	//fmt.Printf("%p \n",node)
	//for node.next != nil {
	//	// 啥也不做 就移动指针
	//	node = node.next
	//
	node := d.tail //直接拿到最后一个 减低时间复杂度
	newNode := NewNode(v)
	newNode.per = node
	node.next = newNode // 一定是next 不然修改的只是指针移动的每个元素项
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
	doubly := createDoublyList()
	doubly.RPush("C++")
	doubly.RPush("Java")
	doubly.RPush("Golang")
	fmt.Println(doubly.length)
	fmt.Printf("头节点是:%p\n", doubly.head)
	fmt.Println("尾节点是:", doubly.tail)
	fmt.Println(doubly.getNode(3).value)
	fmt.Println(doubly.getNode(1).value)
	fmt.Printf("头节点是:%p\n", doubly.head)
	fmt.Println("尾节点是:", doubly.tail)
}
