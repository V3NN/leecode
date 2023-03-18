package __removeDuplicates

/**
26. 删除有序数组中的重复项

给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
将最终结果插入 nums 的前 k 个位置后返回 k 。
不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

*/

func removeDuplicates(nums []int) int {
	// 定义两个指针， idx 表示是遍历的索引， cur表示当前不重复的索引位置。因为nums[0]是第一个出现的元素，不存在重复，所以初始化为1
	cur, idx := 1, 1
	// 开始循环
	for idx < len(nums) {
		// 如果第一个值（index=0)与第二个值（index=1)不相等
		if nums[idx] != nums[idx-1] {
			// 把当前遍历到的值复制到已遍历完成的索引位，同时已遍历完的索引位置+1
			nums[cur] = nums[idx]
			cur++
		}
		// 循环继续
		idx++
	}
	return cur
}
