package __romanToInt

// 定义哈希表
var m = map[byte]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
}

func romanToInt(s string) (ans int) {
	l := len(s) - 1
	for i := range s {
		// 取出字符串第一个字符对应的数值
		val := m[s[i]]
		// 如果当前值小于了下一个字符对应的数值则总值减去当前值，否则相加
		if i < l && val < m[s[i+1]] {
			ans -= val
		} else {
			ans += val
		}
	}

	return ans
}
