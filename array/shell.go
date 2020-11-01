package main

import (
	"fmt"
)

func main() {
	arr := []int{14, 12, 9, 7, 81, 2, 3, 4, 1}
	sortShell(arr)
	fmt.Println(arr)
}

// 希尔排序 O(n^n)
func shell(arr []int) {
	// gap = pace = 步伐
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

func sortShell(arr []int) {
	var index, value int
	// 缩小数据增量
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		// 然后把缩小后的数据看成一个整体
		for i := gap; i < len(arr); i++ {
			// index记录需要插入数的启示坐标,value记录插入的数据
			index, value = i, arr[i]
			// 判断要插入的数据比前面的-index 小
			if arr[i] < arr[i-index] {
				// 依次的判断是否找到对应的插入的位置
				for i-gap >= 0 && value < arr[i-gap] {
					// 让插入数的向后移动
					arr[i] = arr[i-gap]
					// ？？？
					i -= gap
				}
				arr[i] = value
			}
		}
	}
}

func ShellSort(a []int) {
	n := len(a)
	h := 1
	for h < n/3 { //寻找合适的间隔h
		h = 3*h + 1
	}
	for h >= 1 {
		//将数组变为间隔h个元素有序
		for i := h; i < n; i++ {
			//间隔h插入排序
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				swap(a, j, j-h)
			}
		}
		h /= 3
	}
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
