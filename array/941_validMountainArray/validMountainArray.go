package __validMountainArray

/**

  941. 有效的山脉数组

	给定一个整数数组 arr，如果它是有效的山脉数组就返回 true，否则返回 false。

  示例 1：
	输入：arr = [2,1]
	输出：false

  示例 2：
	输入：arr = [3,5,5]
	输出：false

*/

func validMountainArray(arr []int) bool {
	// 左右两侧起始点
	left := 0
	l := len(arr)
	right := l - 1

	// 左侧第一位与第二位比较 如果是递增同时小于右侧末尾则向中间靠拢
	for left+1 < l && arr[left] < arr[left+1] {
		left++
	}
	// 右侧最后一位与倒数第二位比较 如果最后一位小于倒数第二位则向中间靠拢
	for right > 0 && arr[right] < arr[right-1] {
		right--
	}
	// 左右侧是否都指向最高点且不能位于两端
	return left == right && left > 0 && right < l-1
}
