package _142_环形链表_II

type ListNode struct {
	Val int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return nil
	}
	slow, fast := head, head
	for fast!=nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			break
		}
	}
	if slow != fast{
		return nil
	}
	newSlow := head
	for newSlow != fast{
		newSlow = newSlow.Next
		fast = fast.Next
	}
	return newSlow
}
