package __permutations

import "testing"

func TestPermute(t *testing.T) {
	arr := []int{1, 2, 3}
	t.Log(permute(arr))

	t.Log(permute([]int{0, 1}))
}
