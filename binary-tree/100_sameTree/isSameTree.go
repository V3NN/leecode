package __sameTree

/**

  100. 相同的树

	给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

	示例 ：
	输入：p = [1,2,3], q = [1,2,3]
	输出：true
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 如果两个都为空则相同
	if p == nil && q == nil {
		return true
	}
	// 如果只有一个为空则不相同
	if p == nil || q == nil {
		return false
	}
	// 值不相等不相同
	if p.Val != q.Val {
		return false
	}
	// 判断子树是否相等
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
