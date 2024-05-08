package __containsDuplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	t.Log(containsDuplicate([]int{1, 2, 3, 1}))
	t.Log(containsDuplicate([]int{1, 2, 3, 4}))
	t.Log(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
