package main

import (
	"fmt"
)

func main() {
	var arr []int = []int{33, 12, 11, 23, 4, 5, 6, 1}
	fmt.Println(merge(arr))
}
func merge(arr []int) []int {
	var result []int
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2
	left := arr[:middle]
	right := arr[middle:]
	return func(left, right []int) []int {
		l, r := 0, 0
		for len(left) > l && len(right) > r {
			if arr[l] < arr[r] {
				result = append(result, left[l])
				l++
			} else {
				result = append(result, right[r])
				r++
			}
		}
		result = append(result, left[l:]...)
		result = append(result, right[r:]...)
		return result
	}(merge(left), merge(right))
}
