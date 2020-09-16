package main

import "fmt"

// 反正字符串 把一个字符串的内容提供程序颠倒过来
// https://www.geekxh.com/1.3.%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%B3%BB%E5%88%97/301.html#_02%E3%80%81%E9%A2%98%E7%9B%AE%E5%9B%BE%E8%A7%A3
func main() {
	reverseString("您好啊")
}

func reverseString(str string) {
	var r []rune = []rune(str)
	// 定义2个下标 一个从前往后移动  一个从后往前移动  前下标大于后下标说明交换完毕
	for index, later := 0, len(r)-1; later > index; index, later = index+1, later-1 {
		r[index], r[later] = r[later], r[index]
	}
	fmt.Println(string(r))
}
