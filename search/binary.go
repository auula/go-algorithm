package main

import "fmt"

func main() {
	arr := []int{1, 23, 34, 12, 45, 67, 78}
	fmt.Println(binarySearch(12, 0, len(arr), arr))
}

func binarySearch(v, left, right int, arr []int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if arr[mid] == v {
		return mid // index
	}
	if arr[mid] < v {
		return binarySearch(v, mid+1, right, arr)
	}
	if arr[mid] > v {
		return binarySearch(v, left, mid-1, arr)
	}
	return -1
}
