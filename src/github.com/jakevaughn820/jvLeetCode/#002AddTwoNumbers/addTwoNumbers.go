package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	lSum := 0
	for current := ret; l1 != nil || l2 != nil || lSum != 0; current = current.Next {
		if l1 != nil {
			lSum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			lSum += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: lSum % 10}
		lSum /= 10
	}
	return ret.Next
}
