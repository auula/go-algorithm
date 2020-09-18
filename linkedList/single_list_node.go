package main

import "fmt"

// go单向链表实现
// https://media.geeksforgeeks.org/wp-content/uploads/20200318172830/ezgif.com-gif-maker2.gif
type singleNode struct {
	Value    int         `json:"value"` // 节点保存的值
	NextNode *singleNode //每个节点里面又存储了下一个节点的地址
}

// 头 尾 长度
type linkedList struct {
	Head   *singleNode
	Length uint64
}

/*
   当链表中只包含头节点，则链表长度为0，称为空链表。 当插入元素是，也是忽略头节点进行的操作。
*/

// 创建空链表
//创建链表的时候，不需要传入参数，直接返回一个链表就ok
func create() *linkedList {
	return &linkedList{
		Head:   new(singleNode),
		Length: 0,
	}
}

func (l *linkedList) isEmpty() bool {
	return l.Head.NextNode == nil
}

func (l *linkedList) append(v int) {
	node := l.Head             //取得头部节点
	for node.NextNode != nil { // 循环移动指针到最后一个节点
		node = node.NextNode // 反复取下一个节点
	}
	// 如果是nil说明没有下一个直接塞进去node
	node.NextNode = &singleNode{Value: v, NextNode: nil}
	l.Length++ //加加
}

func (l *linkedList) remove(index int) {
	node := l.Head // 拿到头部节点的指针
	// 万变不离其中 就是操作链表指针到达要移除元素的下标前一个位置
	for i := 0; i < index-1; i++ {
		node = node.NextNode // 循环一次移动一次指针
	}
	node.NextNode = node.NextNode.NextNode // 把节点的next指向到下一个的下一个next 从而达到移除的效果
	l.Length--
}

func (l *linkedList) allElement() {
	for node := l.Head; node != nil; node = node.NextNode {
		fmt.Print(node.Value, "->")
	}
	fmt.Println()
}

func (l *linkedList) get(index int) int {
	node := l.Head
	for i := 0; i < index; i++ {
		node = node.NextNode
	}
	return node.Value
}

func (l *linkedList) update(index, value int) {
	node := l.Head
	for i := 0; i < index; i++ {
		node = node.NextNode
	}
	node.Value = value
}

func (l *linkedList) insert(index, value int) {
	node := l.Head
	// 减一是因为指针移动是要插入的位置 如果要插入到这个位置所有就是修改前面的nextNode
	for i := 0; i < index-1; i++ {
		node = node.NextNode
	}
	temp := node.NextNode
	node.NextNode = &singleNode{Value: value, NextNode: temp}
	l.Length++
}

func main() {
	list := create()
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.append(5)
	fmt.Println(list.Length)
	list.allElement()
	list.remove(4)
	list.allElement()
	fmt.Println(list.Length)
	fmt.Println(list.get(4))
	list.update(4, 66)
	fmt.Println(list.get(4))
	list.insert(3, 22)
	list.allElement()
	fmt.Println(list.Length)
}

//5
//0->1->2->3->4->5->
//0->1->2->3->5->
//4
//5
//66
//0->1->22->2->3->66->
