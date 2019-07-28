package _021_合并两个有序链表

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 有一条链为nil，直接返回另一条链
	if l1 == nil{
		return l2
	}

	if l2 == nil{
		return l1
	}

	// 此时，两条链都不为nil，可以直接使用l.Val，而不用担心panic
	// 在l1和l2之间，选择较小的节点作为head，并设置好node
	var head, node *ListNode
	if l1.Val < l2.Val{
		head = l1
		node = l1
		l1 = l1.Next
	}else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			node.Next = l1
			l1 = l1.Next
		}else {
			node.Next = l2
			l2 = l2.Next
		}
		// 有了这一步，head才是一个完整的链
		node = node.Next
	}

	if l1 != nil{
		node.Next = l1
	}
	if l2 != nil{
		node.Next = l2
	}
	return head

}
