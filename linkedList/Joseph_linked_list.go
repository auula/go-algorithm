package main

import (
	"fmt"
)

// 循环链表解决约瑟夫问题

type JosephNode struct {
	value int
	next  *JosephNode
}

type JosepList struct {
	head   *JosephNode
	length int
}

func generate() *JosepList {
	var new *JosephNode
	var first *JosephNode
	var node *JosephNode
	for i := 1; i <= 41; i++ {
		if i == 1 {
			// 是不是第一个节点
			first = &JosephNode{value: i, next: new}
			node = first // 第一个节点
			continue
		}
		new = &JosephNode{value: i, next: nil}
		node.next = new // 让上一组的节点指向这个节点
		node = new      // 更新一下暂存的节点
		if i == 41 {
			node.next = first // 设置成环表
		}
	}
	return &JosepList{head: &JosephNode{value: 0, next: first}, length: 41}
}

func (jl *JosepList) allPrint() {
	for node := jl.head.next; node != nil; node = node.next {
		fmt.Print(node.value, "->")
		if node.value == 41 {
			// 因为是环表
			break
		}
	}
	fmt.Println()
}

func (jl *JosepList) suicide() {
	count := 0
	var before *JosephNode
	node := jl.head.next
	for node != node.next { // 因为是环表 尾巴指向头 然后直到最后自己指向自己结束
		count++
		if count == 3 {
			before.next = node.next
			fmt.Print(node.value, ",")
			count = 0
			node = node.next
		} else {
			before = node // 暂存前一个节点 方便后面遇到3使用
			node = node.next
		}
	}
	fmt.Println(node.value)
}

func main() {
	josepList := generate()
	josepList.allPrint()
	josepList.suicide()
}
