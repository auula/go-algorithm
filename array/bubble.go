// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/10/15 - 3:29 下午 - UTC/GMT+08:00

package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 9, -1, 10, 20}
	//bubbleSort(arr)
	sort(arr)
	fmt.Println(arr)
}

func Bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		// 冒泡是每次相邻的2个元素排
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func sort(arr []int) {
	var flag bool
	for i := 0; i < len(arr); i++ {
		// 冒泡是每次相邻的2个元素排 减去已经排好序的元素个数 i其实就是排好序的个数
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		// 优化一次就排好的序的数据情况
		if !flag {
			break
		} else {
			flag = false
		}
	}

}
