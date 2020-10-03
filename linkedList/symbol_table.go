// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/10/3 - 5:31 下午 - UTC/GMT+08:00

package main

import "fmt"

// 符号表实现
type SymbolNode struct {
	Key, Value interface{}
	Next       *SymbolNode
}

type SymbolTable struct {
	Head *SymbolNode
	Size int
}

func NewSymbolTable() *SymbolTable {
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
	var prevNode *SymbolNode
	node := t.Head
	for node != nil {
		prevNode = node
		node = node.Next
		if node.Key == key {

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
	table := NewSymbolTable()
	table.put("k1", "v1")
	fmt.Println(table.get("k1"))
	table.put("k1", "v11")
	fmt.Println(table.get("k1"))
}
