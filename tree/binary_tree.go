// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/10/5 - 9:43 下午 - UTC/GMT+08:00


// 二叉树 查询树设计实现
package main

// 树的节点
type treeNode struct {
	Key int
	Value interface{}
	Left,Right *treeNode
}

func newTreeNode(key int,value interface{},node ...*treeNode) *treeNode {
	//fmt.Printf("%T \n",node)//[]*main.treeNode
	//fmt.Println(node[0])
	return &treeNode{Key: key,Value: value}
}

func main() {
	newTreeNode(1,"v1",nil,nil)
}

type BinaryTree struct {
	Root *treeNode
	Length int
}