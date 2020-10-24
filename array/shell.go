package main

import (
	"fmt"
)

func main() {
	arr := []int{14, 12, 9, 7, 81, 2, 3, 4, 1}
	shell(arr)
	fmt.Println(arr)
}

// 希尔排序 O(n^n)
func shell(arr []int) {
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			for j := i - gap; j >= 0; j -= gap {
				if arr[j] > arr[j+gap] {
					arr[j], arr[j+gap] = arr[j+gap], arr[j]
				}
			}
		}
	}
}
