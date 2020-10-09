// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/10/5 - 9:43 下午 - UTC/GMT+08:00

// 二叉树 查询树设计实现
package main

import "fmt"

// 树的节点
type treeNode struct {
	Key         int
	Value       interface{}
	Left, Right *treeNode
}

func newTreeNode(key int, value interface{}, node ...*treeNode) *treeNode {
	//fmt.Printf("%T \n",node)//[]*main.treeNode
	//fmt.Println(node[0])
	return &treeNode{Key: key, Value: value}
}

func (n *treeNode) preOrder() {
	fmt.Println(n.Key) // 前序遍历
	if n.Left != nil {
		n.Left.preOrder()
	}
	if n.Right != nil {
		n.Right.preOrder()
	}
}

func (n *treeNode) infixOrder() {
	if n.Left != nil {
		n.Left.infixOrder()
	}
	fmt.Println(n.Key) // 中序遍历
	if n.Right != nil {
		n.Right.infixOrder()
	}
}
func (n *treeNode) afterOrder() {
	if n.Left != nil {
		n.Left.afterOrder()
	}
	if n.Right != nil {
		n.Right.afterOrder()
	}
	fmt.Println(n.Key) // 后序遍历
}

func (n *treeNode) setLeft(node *treeNode) {
	if n.Left.Key < node.Key {
		n.Left = node
	}
	n.Left.setLeft(node)
}

func (n *treeNode) setRight(node *treeNode) {
	if n.Right.Key < node.Key {
		n.Right = node
	}
	n.Right.setRight(node)
}

func main() {
	bt := new(BinaryTree)
	root := newTreeNode(1, "v1", nil)
	n2 := newTreeNode(2, "v2", nil)
	n3 := newTreeNode(3, "v3", nil)
	n4 := newTreeNode(4, "v4", nil)
	n5 := newTreeNode(5, "v5", nil)
	root.Left = n3
	root.Right = n4
	n5.Left = n2
	n4.Left = n5
	bt.setRoot(root)
	bt.preOrder() // 1,3,4,5,2
	fmt.Println("中序")
	bt.infixOrder() // 3,1,2,5,4
	fmt.Println("后序")
	bt.afterOrder() // 3,2,5,4,1
}

// BinaryTree 二叉树
type BinaryTree struct {
	Root   *treeNode
	Length int
}

func (bt *BinaryTree) setRoot(root *treeNode) {
	bt.Root = root
	//bt.Length++
}

// 前序遍历
func (bt *BinaryTree) preOrder() {
	if bt.Root != nil {
		bt.Root.preOrder()
	}
}

// 中序遍历
func (bt *BinaryTree) infixOrder() {
	if bt.Root != nil {
		bt.Root.infixOrder()
	}
}

// 后序遍历
func (bt *BinaryTree) afterOrder() {
	if bt.Root != nil {
		bt.Root.afterOrder()
	}
}

func (bt *BinaryTree) add(node *treeNode) {
	if node.Key > bt.Root.Key {
		bt.add(node)
	}
}
