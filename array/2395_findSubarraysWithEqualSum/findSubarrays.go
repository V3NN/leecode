package __findSubarraysWithEqualSum

/**
  2395. 和相等的子数组

  给你一个下标从 0 开始的整数数组 nums ，判断是否存在 两个 长度为 2 的子数组且它们的 和 相等。注意，这两个子数组起始位置的下标必须 不相同 。
  如果这样的子数组存在，请返回 true，否则返回 false 。 子数组 是一个数组中一段连续非空的元素组成的序列。

  示例 1：
  输入：nums = [4,2,4]
  输出：true
  解释：元素为 [4,2] 和 [2,4] 的子数组有相同的和 6 。

*/

func findSubarrays(nums []int) bool {
	hash := make(map[int]bool)

	for i := 1; i < len(nums); i++ {
		val := nums[i] + nums[i-1]
		if hash[val] {
			return true
		}
		hash[val] = true
	}

	return false
}
