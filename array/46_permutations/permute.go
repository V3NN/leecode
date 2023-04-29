package __permutations

/**
  46. 全排列

  给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/

func permute(nums []int) [][]int {
	l := len(nums)
	res := make([][]int, 0)
	path := []int{}
	used := make(map[int]bool)

	var dfs func()
	dfs = func() {
		if len(path) == l {
			tmp := make([]int, l)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < l; i++ {
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			path = append(path, nums[i])
			dfs()
			used[nums[i]] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return res
}
