package __palindromicSubstrings

/**
  647. 回文子串

  给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
  回文字符串 是正着读和倒过来读一样的字符串。
  子字符串 是字符串中的由连续字符组成的一个序列。
  具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

  示例 1：
    输入：s = "abc"
    输出：3
    解释：三个回文子串: "a", "b", "c"

*/

func countSubstrings(s string) int {
	l := len(s)
	n := 0
	f := func(left, right int) {
		for left >= 0 && right < l && s[left] == s[right] {
			n++
			left--
			right++
		}
	}

	for i := 0; i < l; i++ {
		f(i, i)
		f(i, i+1)
	}
	return n
}
