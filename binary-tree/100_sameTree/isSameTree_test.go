package __sameTree

import "testing"

func TestIsSameTree(t *testing.T) {
	t1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	t2 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}

	t.Log(isSameTree(t1, t2))

	t3 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	t4 := &TreeNode{1, nil, &TreeNode{3, nil, nil},}

	t.Log(isSameTree(t3, t4))

}