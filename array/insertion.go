// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/10/16 - 4:42 下午 - UTC/GMT+08:00

package main

import "fmt"

func main() {
	arr := []int{11, 12, 13, 2, 49, 11, 4, 6, 3, 1}
	insertionSort(arr)
	//selection(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) {
	var insertVal, index int
	for i := 0; i < len(arr)-1; i++ {
		// 把数据一分为二 有序数据为0号下标元素
		// 无序数据为0号下标的元素之后的元素
		// 然后从无序列表中拿第一个元素依次和有序列表的元素比较 从后往前
		insertVal, index = arr[i+1], i
		for index >= 0 && insertVal < arr[index] {
			arr[index+1] = arr[index] // 后移动一位和打扑克一样
			index--                   // 因为刚刚加一了所有减回来记录当前下标
			arr[index+1] = insertVal  // 加一是因为遍历到有序类别0号元素减一是-1 所有加回来
		}
	}
}

func selection(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		mid := i
		for j := i + 1; j < len(arr); j++ {
			if arr[mid] > arr[j] {
				arr[mid], arr[j] = arr[j], arr[mid]
			}
		}
	}
}
