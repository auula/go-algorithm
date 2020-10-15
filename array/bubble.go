// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/10/15 - 3:29 下午 - UTC/GMT+08:00

package main

import "fmt"

func main() {
	arr := []int{6, 5, 4, 3, 2, 1}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	// 每轮加去之前也就排序好的一个元素位置
	// len - 1 是因为
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
