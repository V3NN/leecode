package __maxSubArray

import "testing"

func TestMaxSubArray(t *testing.T) {
	t.Logf("[1,2,-1,3,4,5] MaxSubArray = %d", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	t.Logf("[1] MaxSubArray = %d", maxSubArray([]int{1}))
	t.Logf("[1,2,3,-1,-2,-3,1,2,4] MaxSubArray = %d", maxSubArray([]int{1, 2, 3, -1, -2, -3, 1, 2, 4}))
}
