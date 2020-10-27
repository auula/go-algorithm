package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 1, 2, 55, 6, 123, 2, 7}
	bubble(arr)
	fmt.Println(arr)
	arr = []int{1, 2, 1, 2, 55, 6, 123, 2, 7}
	insertion(arr)
	fmt.Println(arr)
	arr = []int{1, 2, 1, 2, 55, 6, 123, 2, 7}
	selection(arr)
	fmt.Println(arr)
	arr = []int{1, 2, 1, 2, 55, 6, 123, 2, 7}
	shell(arr)
	fmt.Println(arr)
}

func bubble(rows []int) {
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows)-1-i; j++ {
			if rows[j] > rows[j+1] {
				rows[j], rows[j+1] = rows[j+1], rows[j]
			}
		}
	}
}
func insertion(rows []int) {
	var value, index int
	for i := 0; i < len(rows)-1; i++ {
		value, index = rows[i+1], i
		for index >= 0 && value < rows[index] {
			rows[index+1] = rows[index]
			index--
		}
		rows[index+1] = value
	}
}
func selection(rows []int) {
	for i := 0; i < len(rows)-1; i++ {
		mid := i
		for j := i + 1; j < len(rows); j++ {
			if rows[mid] > rows[j] {
				rows[mid], rows[j] = rows[j], rows[mid]
			}
		}
	}
}
func shell(rows []int) {
	for pace := len(rows) / 2; pace > 0; pace /= 2 {
		for i := pace; i < len(rows); i++ {
			for j := i - pace; j >= 0; j -= pace {
				if rows[j] > rows[j+pace] {
					rows[j+pace], rows[j] = rows[j], rows[j+pace]
				}
			}
		}
	}
}
