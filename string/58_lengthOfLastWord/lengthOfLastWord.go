package __lengthOfLastWord

func lengthOfLastWord(s string) int {
	l := 0
	// 倒序遍历
	for i := len(s) - 1; i >= 0; i-- {
		// 如果尾部不为空则长度+1
		if string(s[i]) != " " {
			l++
		}
		// 长度大于0说明已经遍历到字符了，如果当前索引位为空则直接停止循环
		if l > 0 && string(s[i]) == " " {
			break
		}
	}
	return l
}
