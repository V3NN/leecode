package __removeElement

import "testing"

func TestRemoveElement(t *testing.T) {
	arr, val := []int{3, 2, 2, 3}, 3
	t.Logf("[3,2,2,3] remover %d, len = %d", val, removeElement(arr, val))

	arr, val = []int{0, 1, 2, 2, 3, 0, 4, 2}, 2
	t.Logf("[0, 1, 2, 2, 3, 0, 4, 2] remover %d, len = %d", val, removeElement(arr, val))
}
