package __lengthOfLongestSubstring

import "strings"

/**
3. 无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}
	// 第一个字符肯定不重复，默认给第一个
	sub := s[0:1]
	// 默认长度为1
	n := 1
	for i := 1; i < len(s); i++ {
		// 如果当前字符出现在子串中出现则截取出现位置之后的字符
		if idx := strings.Index(sub, string(s[i])); idx != -1 {
			sub = sub[idx+1:]
		}
		// 追加
		sub += string(s[i])
		// 判断长度是否大于前一次的长度
		if n < len(sub) {
			n = len(sub)
		}
	}
	return n
}
