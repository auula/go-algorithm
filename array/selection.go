// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/10/15 - 8:19 下午 - UTC/GMT+08:00

// 选择排序
package main

import "fmt"

func main() {
	arr := []int{6, 5, 4, 3, 2, 1}
	sortSelection(arr)
	fmt.Println(arr)
}

// 选择排序 O(n^2)
func Selection(arr []int) {
	// 这里的减一是因为需要通过下标方法元素 元素下标是从0开始的
	// 假设mid是最小值，然后和后面元素进行比较
	// 如果发现有比mid小的元素，就更新mid坐标
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		// 一轮结束之后交换2个元素位置
		arr[i], arr[min] = arr[min], arr[i]
	}
}
