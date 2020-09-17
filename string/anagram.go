package main

import (
	"fmt"
)

// 题目要求:
// 什么是异位词? 字母异位词指字母相同，但排列不同的字符串。
// 解题思路
//创建一个数组做26个字母的记数器，遍历一遍字符串，若s中存在某，则计数器相应的字母数加1，同理t中为减1，最后若计数器各个元素均为0，即返回true，否则false.
//作者：wang-xiao-zhu-3
//链接：https://leetcode-cn.com/problems/valid-anagram/solution/242-you-xiao-de-zi-mu-yi-wei-ci-goyu-yan-jie-fa-by/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	fmt.Println(isAnagram("anagram", "agramna"))
}

// unicode编码 https://pic3.zhimg.com/v2-11f9b8b92020bf7e54ccb800750bf902_r.jpg
func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	counter := [26]int{} // 底层模拟出来了一个26字母表
	for i := 0; i < len(str1); i++ {
		// 把第一个字符串里面的字符出现的标记位加1
		counter[str1[i]-97]++
		// 把第二个字符串里面的字符出现的标记减1 这样就达到了检测2个字符串的字符是否都一样效果
		// unicode编码
		counter[str2[i]-97]--
	}
	for _, j := range counter {
		if j != 0 {
			return false
		}
	}
	return true
}
