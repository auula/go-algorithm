// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/10/16 - 4:42 下午 - UTC/GMT+08:00

package main

import "fmt"

func main() {
	arr := []int{12, 13, 2, 49, 11, 4, 6, 3, 1}
	//insertionSort(arr)
	selection(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				break
			}
		} 
	}
}

func selection(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		mid := i
		for j := i+1; j < len(arr); j++ {
			if arr[mid] > arr[j] {
				arr[mid],arr[j] = arr[j],arr[mid]
			}
		}
	}
}