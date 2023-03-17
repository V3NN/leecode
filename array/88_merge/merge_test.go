package __merge

import "testing"

func TestMerge(t *testing.T) {
	nums1, m := []int{1, 2, 3, 0, 0, 0}, 3
	nums2, n := []int{2, 5, 6}, 3
	src := make([]int, len(nums1))
	copy(src, nums1)
	merge(nums1, m, nums2, n)
	t.Logf("%v merge %v = %v", src, nums2, nums1)

	nums1, m = []int{0}, 0
	nums2, n = []int{1}, 1
	src = make([]int, len(nums1))
	copy(src, nums1)
	merge(nums1, m, nums2, n)
	t.Logf("%v merge %v = %v", src, nums2, nums1)
}
