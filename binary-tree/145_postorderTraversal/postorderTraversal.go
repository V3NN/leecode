package __postorderTraversal

/**
145. 二叉树的后序遍历

给你二叉树的根节点 root ，返回它节点值的 后序 遍历。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var f func(n *TreeNode)
	f = func(n *TreeNode) {
		if n == nil {
			return
		}
		// 后序遍历 左-右-中
		f(n.Left)
		// 右节点
		f(n.Right)
		// 中(root)节点
		res = append(res, n.Val)
	}
	f(root)
	return res
}
