package __combinationSum2

/**

40. 组合总和 II

  给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
  candidates 中的每个数字在每个组合中只能使用 一次 。
  注意：解集不能包含重复的组合。

  输入: candidates = [10,1,2,7,6,1,5], target = 8,
  输出:
	[
		[1,1,6],
		[1,2,5],
		[1,7],
		[2,6]
	]

*/

import (
	"sort"
)

var (
	res  [][]int
	path []int
)

func combinationSum2(candidates []int, target int) [][]int {
	res, path = make([][]int, 0), make([]int, 0)
	sort.Ints(candidates)
	helper(candidates, target, 0)
	return res
}

func helper(candidates []int, target, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		// 早到组合了
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		// 因为数组排好序了，小的数字都得不到结果，大的数字就没有必要计算了
		if candidates[i] > target {
			break
		}
		// 去重
		if i != start && candidates[i] == candidates[i-1] {
			continue
		}

		path = append(path, candidates[i])
		helper(candidates, target-candidates[i], i+1)
		path = path[:len(path)-1]
	}
}
