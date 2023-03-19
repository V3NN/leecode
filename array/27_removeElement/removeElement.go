package __removeElement

/**
27. 移除元素

给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

*/

func removeElement(nums []int, val int) int {
	// 定义一个指针记录长度
	cur := 0
	for i := 0; i < len(nums); i++ {
		// 如果遍历的元素不是要删除的元素则长度+1并且覆盖当前索引位
		if nums[i] != val {
			nums[cur] = nums[i]
			cur++
		}
	}
	return cur
}
