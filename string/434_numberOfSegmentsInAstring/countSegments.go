package __numberOfSegmentsInAstring

import "strings"

/**

 434. 字符串中的单词数

 统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
 请注意，你可以假定字符串里不包括任何不可打印的字符。


 输入: "Hello, my name is John"
 输出: 5
解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。

*/

// 方法一
// 直接调用库函数
func countSegments(s string) int {
	return len(strings.Fields(s))
}

// 方法二
// 遍历字符串，若当前下标之前为空格（或者为初始下标），且自身不为空格，则其为单词开始的下标

//func countSegments(s string) int {
//	count := 0
//	for i := 0; i < len(s); i++ {
//		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
//			count++
//		}
//	}
//
//	return count
//}
