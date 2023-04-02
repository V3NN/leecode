package __mergeTwoSortedLists

/**

21. 合并两个有序链表

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := new(ListNode)
	node := res
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			node.Next = list2
			list2 = list2.Next
		} else {
			node.Next = list1
			list1 = list1.Next
		}
		node = node.Next
	}
	if list1 != nil {
		node.Next = list1
	}
	if list2 != nil {
		node.Next = list2
	}
	return res.Next
}
