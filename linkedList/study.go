package main

import "fmt"

type Object interface{}

//节点
type Node struct {
	data Object
	next *Node
}

//单向链表
type List struct {
	head *Node
	tail *Node
	size uint64
}

//初始化
func (list *List) Init() {
	(*list).size = 0   // 此时链表是空的
	(*list).head = nil // 没有头
	(*list).tail = nil // 没有尾
}

//向尾部添加数据
func (list *List) Append(node *Node) bool {
	if node == nil {
		return false
	}

	(*node).next = nil
	// 将新元素放入单链表中
	if (*list).size == 0 {
		(*list).head = node
	} else {
		oldTail := (*list).tail
		(*oldTail).next = node
	}

	// 调整尾部位置，及链表元素数量
	(*list).tail = node // node成为新的尾部
	(*list).size++      // 元素数量增加

	return true
}

//插入数据
func (list *List) Insert(i uint64, node *Node) bool {
	// 空的节点、索引超出范围和空链表都无法做插入操作
	if node == nil || i > (*list).size || (*list).size == 0 {
		return false
	}

	if i == 0 { // 直接排第一，也就领导小舅子才可以
		(*node).next = (*list).head
		(*list).head = node
	} else {
		// 找到前一个元素
		preItem := (*list).head
		for j := 1; uint64(j) < i; j++ { // 数前面i个元素
			preItem = (*preItem).next
		}
		// 原来元素放到新元素后面,新元素放到前一个元素后面
		(*node).next = (*preItem).next
		(*preItem).next = preItem
	}

	(*list).size++

	return true
}

//删除元素
func (list *List) Remove(i uint64, node *Node) bool {
	if i >= (*list).size {
		return false
	}

	if i == 0 { // 删除头部
		node = (*list).head
		(*list).head = (*node).next
		if (*list).size == 1 { // 如果只有一个元素，那尾部也要调整
			(*list).tail = nil
		}
	} else {
		preItem := (*list).head
		for j := 1; uint64(j) < i; j++ {
			preItem = (*preItem).next
		}

		node = (*preItem).next
		(*preItem).next = (*node).next

		if i == ((*list).size - 1) { // 若删除的尾部，尾部指针需要调整
			(*list).tail = preItem
		}
	}
	(*list).size--
	return true
}

//获取元素
func (list *List) Get(i uint64) *Node {
	if i >= (*list).size {
		return nil
	}

	item := (*list).head
	for j := 0; uint64(j) < i; j++ { // 从head数i个
		item = (*item).next
	}

	return item
}
func main() {
	var list = List{}
	list.Init()
	for i := 1; i < 100; i++ {
		var node = Node{data: i}
		list.Append(&node)
	}
	var node = list.Get(35)
	fmt.Printf("当前节点位置：%+q \n,数据：%d \n", node, node.data)
	var deleteNode = &Node{}
	var result = list.Remove(35, deleteNode)
	fmt.Printf("删除是否成功%+q \n", result)

	var node2 = list.Get(35)
	fmt.Printf("当前节点位置：%+q \n,数据：%d \n", node2, node2.data)
}
