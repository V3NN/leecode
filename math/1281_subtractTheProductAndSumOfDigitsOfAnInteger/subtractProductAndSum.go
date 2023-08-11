package __subtractTheProductAndSumOfDigitsOfAnInteger

/**

  1281. 整数的各位积和之差
  给你一个整数 n，请你帮忙计算并返回该整数「各位数字之积」与「各位数字之和」的差。

  输入：n = 234
  输出：15
  解释：
	各位数之积 = 2 * 3 * 4 = 24
	各位数之和 = 2 + 3 + 4 = 9
	结果 = 24 - 9 = 15

*/

func subtractProductAndSum(n int) int {
	x, y := 1, 0
	for n > 0 {
		i := n % 10
		n /= 10

		x *= i
		y += i
	}
	return x - y
}
