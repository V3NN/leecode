package __longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	// 使用数组第一个字符串作为对比
	target := strs[0]
	for i := 0; i < len(target); i++ {
		// 循环数组中剩余的字符串
		for j := 1; j < len(strs); j++ {
			// 如果基准字符串长度大于或等于了对比的字符串长度或者字符不相等则截取基准字符串返回
			if i >= len(strs[j]) || target[i] != strs[j][i] {
				return target[:i]
			}
		}
	}
	// 全部相等则直接返回
	return target

}
