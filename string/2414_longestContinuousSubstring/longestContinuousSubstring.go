package __longestContinuousSubstring

/**

2414. 最长的字母序连续子字符串的长度

字母序连续字符串 是由字母表中连续字母组成的字符串。换句话说，字符串 "abcdefghijklmnopqrstuvwxyz" 的任意子字符串都是 字母序连续字符串 。

  例如，"abc" 是一个字母序连续字符串，而 "acb" 和 "za" 不是。
  给你一个仅由小写英文字母组成的字符串 s ，返回其 最长 的 字母序连续子字符串 的长度。

示例 1：
 输入：s = "abacaba"
 输出：2
 解释：共有 4 个不同的字母序连续子字符串 "a"、"b"、"c" 和 "ab" 。 "ab" 是最长的字母序连续子字符串。

*/

func longestContinuousSubstring(s string) int {
	// 储存最大连续长度
	m := 0
	// 存储当前连续长度
	l := 1
	for i := 1; i < len(s); i++ {
		//ASCII值相差1表示是连续的,当前连续长度+1
		if s[i]-s[i-1] == 1 {
			l++
		} else {
			// 不连续的话更新下最大连续长度&当前连续长度复位
			m = max(m, l)
			l = 1
		}
	}
	// 返回最大的连续长度
	return max(m, l)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
