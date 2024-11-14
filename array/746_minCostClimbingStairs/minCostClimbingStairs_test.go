package __minCostClimbingStairs

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	arr := []int{10, 15, 20}
	arr = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	t.Log(minCostClimbingStairs(arr))
}
