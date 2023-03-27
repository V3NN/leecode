package __inorderTraversa

/**
94. 二叉树的中序遍历

给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	// 结果集
	res := make([]int, 0)
	// 定义一个递归体
	var f func(n *TreeNode)
	f = func(n *TreeNode) {
		if n == nil {
			return
		}
		// 中序遍历 左-中-右
		// 左节点
		f(n.Left)
		// 中节点（root)
		res = append(res, n.Val)
		// 右节点
		f(n.Right)
	}
	// 执行递归
	f(root)
	return res
}
