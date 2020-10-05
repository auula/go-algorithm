package main

import "fmt"

// Key 自定义存储类型的key
type Key struct {
	value int
}

// Standard 比较接口 如果要安装key顺序插入就要实现这个接口
type  StandardKey interface {
	CompareTo(key StandardKey) int // 比较函数
	Value() int                    // 获取值方法
}

func (k *Key) CompareTo(key StandardKey) int {
	if k.Value() > key.Value() {
		return 1
	}
	if k.Value() == key.Value() {
		return 0
	}
	return -1
}

func (k *Key) Value() int {
	return k.value
}

// PutByOrderlyKey 通过key排序来插入
func (t *OrderlySymbolTable) PutByOrderlyKey(key StandardKey, value interface{}) {
	// 1.通过key来比较 如果小于就在当前位置插入
	prev := t.Head
	curr := t.Head.Next
	// 循环遍历表 条件是不为空并且key比较的是 当前key小于当前节点的key就继续移动
	for curr != nil && key.CompareTo(curr.K) == 1 {
		prev = curr
		curr = curr.Next
	}
	// 如果当前节点不为空并且比较的key是相等的就直接替换值
	if curr != nil && key.CompareTo(curr.K) == 0 {
		curr.Value = value
		return
	}
	newNode := &OrderlySymbolNode{K: key, Value: value, Next: curr}
	prev.Next = newNode
	t.Size++
}

func (t *OrderlySymbolTable) Get(key StandardKey) *OrderlySymbolNode {
	curr := t.Head
	for curr.Next != nil {
		curr = curr.Next
		if curr.K == key {
			return curr
		}
	}
	return nil
}

// Orderly symbol table
func main() {
	symbolTable := NewOrderlySymbolTable()

	symbolTable.PutByOrderlyKey(&Key{value: 1}, "v1")
	k := &Key{value: 3}
	symbolTable.PutByOrderlyKey(k, "v3")
	symbolTable.PutByOrderlyKey(k, "v33")
	fmt.Println(symbolTable.Size)
	fmt.Println(symbolTable.Get(k))
}

// OrderlySymbolNode is Node
type OrderlySymbolNode struct {
	K     StandardKey
	Value interface{}
	Next  *OrderlySymbolNode
}

// OrderlySymbolTable is Symbol Table
type OrderlySymbolTable struct {
	Head *OrderlySymbolNode
	Size int
}

func NewOrderlySymbolTable() *OrderlySymbolTable {
	return &OrderlySymbolTable{Head: &OrderlySymbolNode{Value: nil, K: nil, Next: nil}, Size: 0}
}
