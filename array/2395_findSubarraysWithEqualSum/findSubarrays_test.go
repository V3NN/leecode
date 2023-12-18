package __findSubarraysWithEqualSum

import "testing"

func TestFindSubarrays(t *testing.T) {
	t.Log(findSubarrays([]int{4, 2, 4}))
	t.Log(findSubarrays([]int{1, 2, 3, 4, 5}))
	t.Log(findSubarrays([]int{0, 0, 0}))
}
