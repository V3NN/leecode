package __searchInsert

import "testing"

func TestSearchInsert(t *testing.T) {
	arr := []int{1, 3, 5, 6}
	target := 5
	t.Logf("%v search %d insert = %d\n", arr, target, searchInsert(arr, target))

	arr = []int{1, 3, 5, 6}
	target = 2
	t.Logf("%v search %d insert = %d\n", arr, target, searchInsert(arr, target))

	arr = []int{1, 3, 5, 6}
	target = 7
	t.Logf("%v search %d insert = %d\n", arr, target, searchInsert(arr, target))

	arr = []int{1}
	target = 7
	t.Logf("%v search %d insert = %d\n", arr, target, searchInsert(arr, target))

}
