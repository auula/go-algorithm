// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/10/15 - 8:19 下午 - UTC/GMT+08:00

// 选择排序
package main

import "fmt"

func main() {
	arr := []int{6, 5, 4, 3, 2, 1}
	selectionSort(arr)
	arr = []int{6, 5, 4, 3, 2, 1}
	sort(arr)
	fmt.Println(arr)
}

// 选择排序 O(n^2) 因为不是每个元素都需要交换 所有比冒泡快
func sort(arr []int) {
	// 这里的减一是因为需要通过下标方法元素 元素下标是从0开始的
	for i := 0; i < len(arr)-1; i++ {
		mid := i
		// 而选择排序是 给定一个元素依次和后面的排序
		// 如果发现逆向就交换这了元素的位置 直到整个走完
		for j := i + 1; j < len(arr); j++ {
			if arr[mid] > arr[j] {
				arr[mid], arr[j] = arr[j], arr[mid]
			}
		}
	}
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
