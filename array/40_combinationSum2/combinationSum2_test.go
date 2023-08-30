package __combinationSum2

import "testing"

func TestCombinationSum2(t *testing.T) {
	t.Log(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 10))
	t.Log(combinationSum2([]int{2, 5, 2, 1, 2, 3, 0}, 5))
	t.Log(combinationSum2([]int{1, 1, 2}, 3))
}
