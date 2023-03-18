package __sortColors

import "testing"

func TestSortColors(t *testing.T) {
	arr := []int{2, 0, 2, 2, 1, 1, 2, 0}
	sortColors(arr)
	t.Logf("[2, 0, 2, 2, 1, 1, 2, 0] sort by colors = %v", arr)

	arr = []int{2, 0, 1}
	sortColors(arr)
	t.Logf("[2,0,1] sort by colors = %v", arr)
}
