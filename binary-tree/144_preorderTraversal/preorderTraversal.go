package __preorderTraversal

/**
144. 二叉树的前序遍历

给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var f func(n *TreeNode)
	f = func(n *TreeNode) {
		if n == nil {
			return
		}
		//前序遍历 中-左-右
		// 中节点（root)
		res = append(res, n.Val)
		// 左节点
		f(n.Left)
		// 右节点
		f(n.Right)
	}
	f(root)
	return res
}
