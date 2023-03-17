package __maxArea

import "testing"

func TestMaxArea(t *testing.T) {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	t.Logf("%v max area = %d", arr, maxArea(arr))

	arr = []int{1, 1}
	t.Logf("%v max area = %d", arr, maxArea(arr))
}
