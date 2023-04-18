package __removeDuplicatesFromSortedList

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	l := &ListNode{Val: 1, Next: &ListNode{1, &ListNode{2, nil}}}
	t.Log(DeleteDuplicates(l))
}
