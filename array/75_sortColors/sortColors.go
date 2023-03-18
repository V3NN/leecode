package __sortColors

/**
75. 颜色分类

给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
必须在不使用库内置的 sort 函数的情况下解决这个问题。
*/

func sortColors(nums []int) {
	// 定义三个指针分别指向数组起始位，遍历位以及结束位
	l, idx, r := 0, 0, len(nums)-1
	// 遍历数组
	for idx <= r {
		// 如果取出的值为0，则交换值
		if nums[idx] < 1 {
			nums[idx], nums[l] = nums[l], nums[idx]
			idx++
			l++
			continue
		}
		// 等于1的时候 因为1在中间位置所以不用变动
		if nums[idx] == 1 {
			idx++
			continue
		}
		// 如果值大于1则将值放置到最后
		if nums[idx] > 1 {
			nums[idx], nums[r] = nums[r], nums[idx]
			r--
		}
	}
}
