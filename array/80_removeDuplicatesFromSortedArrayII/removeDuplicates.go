package __removeDuplicatesFromSortedArrayII

/**

  80. 删除有序数组中的重复项 II

  给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
  不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

*/

func removeDuplicates(nums []int) int {
	l := len(nums)
	// 数组是有序的，保留出现两次的元素，如果长度小于3直接返回
	if l < 3 {
		return l
	}
	// 慢指针表示处理出的数组的长度
	// 快指针表示已经检查过的数组的长度
	slow, fast := 2, 2
	for fast < l {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
