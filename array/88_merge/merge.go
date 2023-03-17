package __merge

/**
/// 88. 合并两个有序数组

给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。
   为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 合并的总长度
	l := m + n - 1
	// 循环两个数组的有效长度
	for ; m > 0 && n > 0; l-- {
		// 如果第一个数组中最后一个值大于第二个数组的最后一个值,则将大的值放到nums1中最后一个位置，同时m-1
		if nums1[m-1] > nums2[n-1] {
			nums1[l] = nums1[m-1]
			m--
		} else {
			// 反之则讲nums2的最后一个值放入nums1最后一个位置，nums2长度-1
			nums1[l] = nums2[n-1]
			n--
		}
	}
	// 如果nums2还没有遍历完（n>0)则将剩余元素逐个填入nums1
	for n > 0 {
		n--
		nums1[n] = nums2[n]
	}

}
