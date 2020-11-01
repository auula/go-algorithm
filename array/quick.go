package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	sortQuick(slice, 0, len(slice)-1)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	// 左右操作下标的指针  一个从头 一个从尾部开始
	left, right := 0, len(a)-1
	// 算出来以那个数为基准
	pivot := rand.Int() % len(a)
	// fmt.Println("pivot=", pivot)
	// fmt.Println("right=", right)
	// fmt.Println("left=", left)
	// 10 19 = 19 10
	// 交换是为下面做准备的  让pivot和right交换
	a[pivot], a[right] = a[right], a[pivot]
	// fmt.Println("a[pivot]", a[pivot])
	// fmt.Println("A =", a)
	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func sortQuick(arr []int, left, right int) {
	var l, r = left, right        // 左右下标
	var pivot = arr[left+right/2] // 中轴值
	for l < r {
		// 在中轴的左边一直找到一个大于或者等于pivot的值
		for arr[l] < pivot {
			l += 1
		}
		// 在中轴的右边一直找到一个小于或者等于pivot的值
		for arr[r] > pivot {
			r -= 1
		}
		arr[r], arr[l] = arr[l], arr[r]
		// 说明左边都是小于pivot的值 右边是大于pivot值
		if l >= r {
			break
		}
		if arr[l] == pivot {
			r -= 1
		}
		if arr[r] == pivot {
			l += 1
		}
		if l == r {
			l++
			r--
		}

	}
	if left < r {
		sortQuick(arr, left, r)
	}
	if right > l {
		sortQuick(arr, l, right)
	}

}
