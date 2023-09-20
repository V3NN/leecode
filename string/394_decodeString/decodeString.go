package __decodeString

import "strconv"

/**
  394. 字符串解码

  给定一个经过编码的字符串，返回它解码后的字符串。
  编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
  你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
  此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

  示例 ：
  输入：s = "3[a]2[bc]"
  输出："aaabcbc"

*/

func decodeString(s string) string {
	res, multi := "", ""
	resStack, multiStack := make([]string, 0, 0), make([]string, 0, 0)

	for i := 0; i < len(s); i++ {
		item := s[i]
		if item <= '9' && item >= '0' {
			multi += string(item)
		} else if item == '[' {
			resStack = append(resStack, res)
			multiStack = append(multiStack, multi)
			res, multi = "", ""
		} else if item == ']' {
			count, _ := strconv.Atoi(multiStack[len(multiStack)-1])
			multiStack = multiStack[:len(multiStack)-1]
			tmp := ""
			for j := 0; j < count; j++ {
				tmp += res
			}
			res = resStack[len(resStack)-1] + tmp
			resStack = resStack[:len(resStack)-1]
		} else {
			res += string(item)
		}
	}
	return res
}
