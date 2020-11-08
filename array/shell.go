package main

import (
	"fmt"
)

func main() {
	row := []int{14, 12, 9, 7, 81, 2, 3, 4, 1}
	// sortShell(arr)
	// fmt.Println(arr)
	Shell(row)
	fmt.Println(row)
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

// 希尔排序 O(n^n)
func Shell(arr []int) {
	var pervIndex, current int
	// 1.先把原数据安装自定义步长分组
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		// 2.然后分好组的数据进行选择排序
		for i := gap; i < len(arr); i++ {
			// 希尔排序
			pervIndex = i - gap
			current = arr[i]
			for pervIndex >= 0 && arr[pervIndex] > current {
				arr[pervIndex+gap] = arr[pervIndex]
				pervIndex -= gap
			}
			if pervIndex+gap != i {
				arr[pervIndex+gap] = current
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

func shellSort(row []int) {
	var increment int
	// 计算增长量 (常用的固定计算增长量的公式)
	for increment < len(row)/2 {
		increment = 2*increment + 1
	}
	// 如果增长量大于等于1说明未完成排序
	for increment >= 1 {
		// i等于每次步伐和间隔 分组的个数 分左右 有序和无限
		for i := increment; i < len(row); i++ {
			// j 等于要传入元素的下标 在增长量里面比较元素
			for j := i; j >= increment; j-- {
				if row[j-increment] > row[j] {
					row[j-increment], row[j] = row[j], row[j-increment]
				} else {
					// 说明找到了合适的位置跳出本次
					break
				}
			}
		}
		// 继续减小增长量
		increment /= 2
	}
}

func simpleShellSort(arr []int) {
	n := len(arr)
	h := 1
	//寻找合适的间隔h
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		// i每轮分组的个数
		for i := h; i < n; i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		h /= 3
	}
}
