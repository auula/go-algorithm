// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/10/3 - 5:31 下午 - UTC/GMT+08:00

package main

import "fmt"

// 符号表实现

// SymbolNode is Node
type SymbolNode struct {
	Key, Value interface{}
	Next       *SymbolNode
}

// SymbolTable is Symbol Table
type SymbolTable struct {
	Head *SymbolNode
	Size int
}

func newSymbolTable() *SymbolTable {
	return &SymbolTable{Head: &SymbolNode{Value: nil, Key: nil, Next: nil}, Size: 0}
}

func (t *SymbolTable) put(key, value interface{}) {
	node := t.Head
	for node.Next != nil {
		// 查找到key 是否存在 如果存在就修改value
		node = node.Next
		if node.Key == key {
			node.Value = value
			return
		}
	}
	// 如果不存在就创建新节点放到头部节点下一个
	firstNode := t.Head.Next
	t.Head.Next = &SymbolNode{Value: value, Key: key, Next: firstNode}
	t.Size++
}

func (t *SymbolTable) remove(key interface{}) {
	node := t.Head
	// 循环遍历表 检测每个节点的key是否匹配
	for node.Next != nil {
		node = node.Next
		if node.Key == key {
			// 如果匹配就删除 删除动作就是吧头部节点下一个节点指向删除节点的下一个节点
			t.Head.Next = node.Next
			t.Size--
			return
		}
	}
}

func (t *SymbolTable) get(key interface{}) *SymbolNode {
	for node := t.Head; node != nil; node = node.Next {
		if node.Key == key {
			return node
		}
	}
	return nil
}

func main() {
	table := newSymbolTable()
	table.put("k1", "v1")
	fmt.Println(table.get("k1"))
	table.put("k1", "v11")
	table.put("k2", 2)
	fmt.Println(table.get("k2"))
	fmt.Println(table.get("k1"))
	table.remove("k2")
	fmt.Println(table.get("k1").Value)
	fmt.Println(table.get("k2"))
}
