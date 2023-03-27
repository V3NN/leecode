package __baseballGame

import (
	"strconv"
)

/**

682. 棒球比赛

你现在是一场采用特殊赛制棒球比赛的记录员。这场比赛由若干回合组成，过去几回合的得分可能会影响以后几回合的得分。
比赛开始时，记录是空白的。你会得到一个记录操作的字符串列表 ops，其中 ops[i] 是你需要记录的第 i 项操作，ops 遵循下述规则：
整数 x - 表示本回合新获得分数 x
"+" - 表示本回合新获得的得分是前两次得分的总和。题目数据保证记录此操作时前面总是存在两个有效的分数。
"D" - 表示本回合新获得的得分是前一次得分的两倍。题目数据保证记录此操作时前面总是存在一个有效的分数。
"C" - 表示前一次得分无效，将其从记录中移除。题目数据保证记录此操作时前面总是存在一个有效的分数。

*/

func calPoints(operations []string) int {
	// 定义stack为解析后的比赛各项分数，count总计分数
	stack, count := make([]int, 0), 0
	for _, v := range operations {
		switch v {
		case "+":
			// 前两项之和
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "C":
			// 删除前一项分数
			stack = stack[:len(stack)-1]
		case "D":
			// 前一项的二倍
			stack = append(stack, stack[len(stack)-1]*2)
		default:
			// 直接添加到栈
			val, _ := strconv.Atoi(v)
			stack = append(stack, val)
		}
	}
	// 遍历栈中元素，累计总分数
	for i := 0; i < len(stack); i++ {
		count += stack[i]
	}

	return count
}
