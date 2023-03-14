package __plusOne

import "testing"

func TestPlusOne(t *testing.T) {
	arr := []int{1, 2, 3}
	t.Log("[]int{1,2,3} = ", plusOne(arr))

	arr = []int{9, 9, 9}
	t.Log("[]int{9,9,9} = ", plusOne(arr))

	arr = []int{1, 9, 9}
	t.Log("[]int{1,9,9} = ", plusOne(arr))
}
