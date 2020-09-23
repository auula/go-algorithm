package main

import "fmt"

// 查询比较多就用顺序表
// 增删比较的多就用链表
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

func (d *doublyList) LPush(v interface{}) {
	firstNode := d.head.next
	newNode := NewNode(v)
	newNode.per = d.head     // 右边插入 都是指向头
	newNode.next = firstNode // 之前的第一个
	d.head.next = newNode
	d.length++
}

func (d *doublyList) insert(index int, v interface{}) {
	node := d.head
	for i := 0; i < index-1; i++ {
		node = node.next
	}
	//afterNode := nil
	newNode := NewNode(v)    //创建新节点
	newNode.per = node       //把新节点的per指向 前一个节点
	newNode.next = node.next // 新节点的next指向 前一个节点的next next就是插入之前的那个节点后面那个节点
	node.next = newNode      // 把旧节点的next指针跟换成下新节点
	d.length++
}

func (d *doublyList) remove(index int) {
	node := d.head
	for i := 0; i < index-1; i++ {
		node = node.next
	}
	if d.tail == node.next { //防止删除尾节点
		return
	}
	next := node.next.next
	next.per = node
	node.next = next
	d.length--
}

func (d *doublyList) getNode(index int) *node {
	node := d.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (d *doublyList) allElement() {
	for node := d.head.next; node != nil; node = node.next {
		fmt.Print(node.value, "->")
	}
	fmt.Println()
}
func main() {
	doubly := createDoublyList()
	doubly.RPush("C++")
	doubly.RPush("Java")
	doubly.RPush("Golang")
	doubly.LPush("C")
	doubly.LPush("B")
	fmt.Println(doubly.length)
	fmt.Printf("头节点是:%p\n", doubly.head)
	fmt.Println("尾节点是:", doubly.tail)
	fmt.Println(doubly.getNode(3).value)
	fmt.Println(doubly.getNode(1).value)
	fmt.Printf("头节点是:%p\n", doubly.head)
	fmt.Println("尾节点是:", doubly.tail)
	doubly.insert(2, "Python")
	fmt.Println(doubly.getNode(4).value)
	fmt.Println(doubly)
	fmt.Println("尾节点是:", doubly.tail)
	//doubly.remove(2) // 删除Java
	fmt.Println(doubly.getNode(2).value)
	fmt.Println(doubly.getNode(1).value)
	doubly.allElement()
}
