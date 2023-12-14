package __checkIfNandItsDoubleExist

/**

  1346. 检查整数及其两倍数是否存在

  给你一个整数数组 arr，请你检查是否存在两个整数 N 和 M，满足 N 是 M 的两倍（即，N = 2 * M）。
  更正式地，检查是否存在两个下标 i 和 j 满足：

  示例 1：
    输入：arr = [10,2,5,3]
    输出：true
  解释：N = 10 是 M = 5 的两倍，即 10 = 2 * 5 。

*/

func checkIfExist(arr []int) bool {
	hash := make(map[int]bool, len(arr))
	for _, val := range arr {
		if ok, _ := hash[val]; ok {
			return true
		}
		hash[val*2] = true

		if val%2 == 0 {
			hash[val/2] = true
		}
	}

	return false
}
