package __longestValidParentheses

/**
  32. 最长有效括号

  给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

  输入：s = "(()"
  输出：2
  解释：最长有效括号子串是 "()"

*/

func longestValidParentheses(s string) int {
	res := 0
	stack := []int{}
	// 栈内存放括号的idx
	stack = append(stack, -1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			// 当遇到右括号时
			// 栈内元素仅有一个，则为上一个右括号位置，需弹出更新
			// 栈内元素多于一个，则为对应左括号位置，需弹出拿下一个值计算
			// 所以无论如何都要弹出
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
