package __reverseVowelsOfAString

import "strings"

/**
345. 反转字符串中的元音字母

给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。

示例 ：
输入：s = "hello"
输出："holle"

*/

func reverseVowels(s string) string {
	vowels := "aAeEiIoOuU"
	l, r := 0, len(s)-1
	res := []byte(s)
	for l < r {
		if !strings.Contains(vowels, string(s[l])) {
			l++
			continue
		}
		if !strings.Contains(vowels, string(s[r])) {
			r--
			continue
		}
		res[l], res[r] = s[r], s[l]
		l++
		r--
	}

	return string(res)
}
