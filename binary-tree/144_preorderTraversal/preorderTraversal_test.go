package __preorderTraversal

import "testing"

func TestPreorderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			2, &TreeNode{3, nil, nil}, nil,
		},
	}
	t.Logf("tree [1,null,2,3] preorder traversal = %v", preorderTraversal(tree))
}
