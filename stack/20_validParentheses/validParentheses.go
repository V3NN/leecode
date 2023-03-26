package __validParentheses

func isValid(s string) bool {
	// 如果字符串长度不为偶数则肯定不符合条件
	if len(s)%2 != 0 {
		return false
	}
	// 定义一个对应关系
	hash := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	// 定义一个切片(栈)
	stack := make([]byte, 0)
	// 循环字符串
	for i := 0; i < len(s); i++ {
		// 如果当前括号再哈希表中出现
		if _, ok := hash[s[i]]; ok {
			// 判断切片长度是否大于0，如过小于或等于0肯定是没有成对的。或者当前出栈的值是否与哈希表中对应
			if len(stack) <= 0 || stack[len(stack)-1] != hash[s[i]] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			// 不存在这写入切片（栈）中
			stack = append(stack, s[i])
		}
	}
	// 如果切片的长度大于0肯定是没有匹配完成的，=0则表示已经匹配完成
	return len(stack) == 0
}
