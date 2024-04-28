package __maximumNumberOfPairsInArray

/**
2341. 数组能形成多少数对

	给你一个下标从 0 开始的整数数组 nums 。在一步操作中，你可以执行以下步骤：

	从 nums 选出 两个 相等的 整数
	从 nums 中移除这两个整数，形成一个 数对
	请你在 nums 上多次执行此操作直到无法继续执行。

	返回一个下标从 0 开始、长度为 2 的整数数组 answer 作为答案，其中 answer[0] 是形成的数对数目，answer[1] 是对 nums 尽可能执行上述操作后剩下的整数数目。

	输入：nums = [1,3,2,1,3,2,2]
	输出：[3,1]
	解释：
	nums[0] 和 nums[3] 形成一个数对，并从 nums 中移除，nums = [3,2,3,2,2] 。
	nums[0] 和 nums[2] 形成一个数对，并从 nums 中移除，nums = [2,2,2] 。
	nums[0] 和 nums[1] 形成一个数对，并从 nums 中移除，nums = [2] 。
	无法形成更多数对。总共形成 3 个数对，nums 中剩下 1 个数字。

*/

func numberOfPairs(nums []int) []int {
	hash := make(map[int]int)
	count := 0
	for num := range nums {
		hash[nums[num]]++
		if hash[nums[num]]%2 == 0 {
			count++
		}
	}
	return []int{count, len(nums) - 2*count}
}

// 遍历一次数组，用一个哈希表保存元素个数的奇偶性，偶数为 false, 奇数则为 true
// 每遇到一个元素，则将奇偶性取反，若取反完后为偶数个，则表明在上次偶数个之后又遇到了两个该元素，可以形成一个数对。
// 最后返回一个数组，第一个元素是数对数，第二个元素是数组长度减去数对数的两倍。
func numberOfPairs2(nums []int) []int {
	count := 0
	hash := map[int]bool{}
	for _, v := range nums {
		if hash[v] {
			count++
		}
		hash[v] = !hash[v]
	}
	return []int{count, len(nums) - count*2}

}
