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
	for node := l.Head.NextNode; node != nil; node = node.NextNode {
		fmt.Print(node.Value, "->")
	}
	fmt.Println()
}

func (l *linkedList) get(index int) *singleNode {
	node := l.Head
	for i := 0; i < index; i++ {
		node = node.NextNode
	}
	return node
}

func (l *linkedList) upData(index, value int) {
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

func (l *linkedList) preAppend(value int) {
	s := &singleNode{Value: value, NextNode: l.Head.NextNode}
	l.Head.NextNode = s // 换到第一个位置
	l.Length++
}

// 返回元素第一次出现的位置
func (l *linkedList) indexOf(value int) uint {
	node := l.Head
	count := 0
	for node != nil {
		if node.Value == value {
			return uint(count)
		}
		count++
		node = node.NextNode
	}
	return 0
}

// 反转条件：
// 1. 想要转换就是要最底3个元素
// 2. 定义一个框
// 3. 每次一点3个元素
func (l *linkedList) reverse() {
	l.Head.NextNode = func(head *singleNode) *singleNode {
		//head.Next, head, prev = prev, head.Next, head
		if head == nil {
			return nil
		}
		var prevNode *singleNode
		var nextNode *singleNode
		curr := head.NextNode
		for curr != nil {
			nextNode = curr.NextNode // 下一个需要交换的元素
			curr.NextNode = prevNode // 反转指针指向前面一个元素
			prevNode = curr          // 下一个需要交换的元素的next指向目前这一个node
			curr = nextNode          // 下次循环的元素指针
		}
		return prevNode
	}(l.Head)
}

// 快慢指针
func (l *linkedList) medianValue() *singleNode {
	slow, fast := l.Head, l.Head
	for fast != nil && fast.NextNode != nil {
		fast = fast.NextNode.NextNode
		slow = slow.NextNode
	}
	return slow
}

// 检查是否有环
func (l *linkedList) isLoop() bool {
	slow, fast := l.Head, l.Head
	for fast != nil && fast.NextNode != nil {
		fast = fast.NextNode.NextNode
		slow = slow.NextNode
		if fast == slow {
			return true
		}
	}
	return false
}

func (l *linkedList) loopEntrance() *singleNode {
	slow, fast := l.Head, l.Head
	var temp *singleNode
	for {
		fast = fast.NextNode.NextNode
		slow = slow.NextNode
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.NextNode == nil {
		return nil
	}
	temp = l.Head
	for temp != slow {
		temp = temp.NextNode
		slow = slow.NextNode
	}
	return slow
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
	list.upData(4, 66)
	fmt.Println(list.get(4))
	list.insert(3, 22)
	list.insert(3, 9999)
	list.preAppend(111)
	list.preAppend(11)
	list.allElement()

	fmt.Println(list.Length)
	fmt.Println("11元素的下标是:", list.indexOf(11))
	fmt.Println(list.Head)
	list.reverse()
	list.allElement()
	fmt.Println(list)
	fmt.Println(list.medianValue().Value)
	list.get(4).NextNode = list.get(3)
	fmt.Println("当前是否为环表:", list.isLoop())
	fmt.Println("当前环的入口是:", list.loopEntrance().Value)
}

//5
//0->1->2->3->4->5->
//0->1->2->3->5->
//4
//5
//66
//0->1->22->2->3->66->
