package __postorderTraversal

import "testing"

func TestPostorderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			2, &TreeNode{3, nil, nil}, nil,
		},
	}
	t.Logf("tree [1,null,2,3] postorder traversal = %v", postorderTraversal(tree))
}
