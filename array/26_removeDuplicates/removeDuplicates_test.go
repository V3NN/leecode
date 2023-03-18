package __removeDuplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Logf("%v Remove Duplicates len's = %d ", arr, removeDuplicates(arr))

	arr = []int{1, 1, 2}
	t.Logf("%v Remove Duplicates len's = %d ", arr, removeDuplicates(arr))
}
