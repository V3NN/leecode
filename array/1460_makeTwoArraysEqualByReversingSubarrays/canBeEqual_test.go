package __makeTwoArraysEqualByReversingSubarrays

import "testing"

func TestCanBeEqual(t *testing.T) {
	t.Log(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
	t.Log(canBeEqual([]int{3, 7, 9}, []int{3, 7, 11}))
}
