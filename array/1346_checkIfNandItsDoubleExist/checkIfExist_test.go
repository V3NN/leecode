package __checkIfNandItsDoubleExist

import "testing"

func TestCheckIfExist(t *testing.T) {
	arr := []int{-2, 0, 10, -19, 4, 6, -8}
	t.Log(checkIfExist(arr))
	arr = []int{0, 0}
	t.Log(checkIfExist(arr))
	arr = []int{7, 1, 14, 11}
	t.Log(checkIfExist(arr))

}
