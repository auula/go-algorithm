package main

import "fmt"

func main() {
	reverseString("Hello Golang")
}

func reverseString(str string) {
	var r []rune = []rune(str)
	// 定义2个下标 一个从前往后移动  一个从后往前移动  前下标大于后下标说明交换完毕
	for index, later := 0, len(r)-1; later > index; index, later = index+1, later-1 {
		r[index], r[later] = r[later], r[index]
	}
	fmt.Println(string(r))
}
