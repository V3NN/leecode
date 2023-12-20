package __makeTwoArraysEqualByReversingSubarrays

/**

  1460. 通过翻转子数组使两个数组相等

  给你两个长度相同的整数数组 target 和 arr 。每一步中，你可以选择 arr 的任意 非空子数组 并将它翻转。
  如果你能让 arr 变得与 target 相同，返回 True；否则，返回 False 。

  示例 1：
  输入：target = [1,2,3,4], arr = [2,4,1,3]
  输出：true
  解释：你可以按照如下步骤使 arr 变成 target：
   1- 翻转子数组 [2,4,1] ，arr 变成 [1,4,2,3]
   2- 翻转子数组 [4,2] ，arr 变成 [1,2,4,3]
   3- 翻转子数组 [4,3] ，arr 变成 [1,2,3,4]
   上述方法并不是唯一的，还存在多种将 arr 变成 target 的方法。

*/

func canBeEqual(target []int, arr []int) bool {
	hash := make(map[int]int, len(target))
	for i := 0; i < len(target); i++ {
		hash[target[i]]++
		hash[arr[i]]--
	}
	for _, val := range hash {
		if val != 0 {
			return false
		}
	}
	return true
}

func canBeEqual2(target []int, arr []int) bool {
	hash := make(map[int]int, len(target))
	for i := 0; i < len(target); i++ {
		hash[target[i]]++
	}
	for i := 0; i < len(arr); i++ {
		hash[arr[i]]--
		if val, ok := hash[arr[i]]; ok && val == 0 {
			delete(hash, arr[i])
		}
	}

	return len(hash) == 0
}
