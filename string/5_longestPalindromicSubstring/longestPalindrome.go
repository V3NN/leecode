package __longestPalindromicSubstring

/**

   最长回文子串

	给你一个字符串 s，找到 s 中最长的回文子串。如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

	示例 1：
	输入：s = "babad"
	输出："bab"
	解释："aba" 同样是符合题意的答案。

*/

func longestPalindrome(s string) string {
	length := 0
	start := 0
	l := len(s)

	expandAroundCenter := func(s string, left, right int) (int, int) {
		// 判断是否回文数，是的话向外扩散
		for left >= 0 && right < l && s[left] == s[right] {
			left--
			right++
		}
		// 如果不是则返回上一个位置
		return left + 1, right - 1
	}

	for i := 0; i < l; i++ {
		// 基数情况判断  aba
		l1, r1 := expandAroundCenter(s, i, i)
		// 偶数情况判断 abba
		l2, r2 := expandAroundCenter(s, i, i+1)

		// 对并当前长度，如果过更大则覆盖保存的值
		if r1-l1 > length-start {
			start, length = l1, r1
		}
		if r2-l2 > length-start {
			start, length = l2, r2
		}
	}

	return s[start : length+1]
}
