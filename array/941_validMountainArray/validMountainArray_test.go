package __validMountainArray

import "testing"

func TestValidMountainArray(t *testing.T) {
	arr := []int{0, 3, 2, 1}
	t.Log(validMountainArray(arr))
}
