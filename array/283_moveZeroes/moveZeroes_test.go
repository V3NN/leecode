package __moveZeroes

import "testing"

func TestMoveZeroes(t *testing.T) {

	arr := []int{0, 1, 0, 3, 12}
	t.Logf("arr before = %v", arr)
	moveZeroes(arr)
	t.Logf("arr after = %v", arr)

	//arr = []int{1, 0, 1}
	//t.Logf("arr before = %v", arr)
	//moveZeroes(arr)
	//t.Logf("arr after = %v", arr)

}
