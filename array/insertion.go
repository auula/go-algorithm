// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/10/16 - 4:42 下午 - UTC/GMT+08:00

package main

import "fmt"

func main() {
	arr := []int{11, 12, 13, 3, 4, 2, 49, 11, 4, 6, 3, 1, -1}
	InsertionSort(arr)
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
		}
		if index+1 != i {
			arr[index+1] = insertVal // 加一是因为遍历到有序类别0号元素减一是-1 所有加回来
		}
	}
}

// InsertionSort 插入排序
func InsertionSort(arr []int) {
	var pervIndex, current int
	for i := 1; i < len(arr); i++ {
		pervIndex = i - 1
		current = arr[i]
		for pervIndex >= 0 && arr[pervIndex] > current {
			arr[pervIndex+1] = arr[pervIndex]
			pervIndex--
		}
		if pervIndex+1 != i {
			arr[pervIndex+1] = current
		}
	}
}

func InsertSort(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}

	for i := 1; i < length; i++ {
		tmp := values[i] // 从第二个数开始，从左向右依次取数
		key := i - 1     // 下标从0开始，依次从左向右

		// 每次取到的数都跟左侧的数做比较，如果左侧的数比取的数大，就将左侧的数右移一位，直至左侧没有数字比取的数大为止
		for key >= 0 && tmp < values[key] {
			values[key+1] = values[key]
			key--
			//fmt.Println(values)
		}

		// 将取到的数插入到不小于左侧数的位置
		if key+1 != i {
			values[key+1] = tmp
		}
		//fmt.Println(values)
	}
}

// ？？？ 个人感觉像是冒泡排序
func simpleInsertionSort(row []int) {
	for i := 1; i < len(row); i++ {
		for j := i; j > 0; j-- {
			if row[j] < row[j-1] {
				row[j], row[j-1] = row[j-1], row[j]
			}
		}
	}
}
func insertSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
