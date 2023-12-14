package __faultyKeyboard

/**

  2810. 故障键盘

   你的笔记本键盘存在故障，每当你在上面输入字符 'i' 时，它会反转你所写的字符串。而输入其他字符则可以正常工作。
   给你一个下标从 0 开始的字符串 s ，请你用故障键盘依次输入每个字符，返回最终笔记本屏幕上输出的字符串。

  示例 1：
	输入：s = "string"
	输出："rtsng"

*/

func finalString(s string) string {
	res := ""
	for k := 0; k < len(s); k++ {
		if s[k] == 'i' {
			res = func(res string) string {
				runes := []rune(res)
				for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
					runes[i], runes[j] = runes[j], runes[i]
				}
				return string(runes)
			}(res)
		} else {
			res += string(s[k])
		}
	}
	return res
}
