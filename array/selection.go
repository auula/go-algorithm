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
	fmt.Println(arr)
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
