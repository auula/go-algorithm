package main

import "fmt"

func main() {
	arr := []int{3, 9, -1, 10, 20}
	insertion(arr)
	fmt.Println(arr)
}

func bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func selection(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		mid := i
		for j := i + 1; j < len(arr); j++ {
			if arr[mid] > arr[j] {
				arr[mid], arr[j] = arr[j], arr[mid]
			}
		}
	}
}

func insertion(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		var value, index = arr[i+1], i
		for index >= 0 && value < arr[i] {
			arr[index+1] = arr[index]
			index--
		}
		arr[index+1] = value
	}
}
