package __binarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{-1, 0, 3, 5, 9, 12}
	target := 2

	t.Logf("%v search [%d] : %d", arr, target, search(arr, target))

}
