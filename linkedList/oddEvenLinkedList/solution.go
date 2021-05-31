package oddEvenLinkedList

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil || head.Next.Next == nil {
		return head
	}
	odd := head
	even := head.Next

	tmpo := odd
	tmpe := even
	for tmpe != nil && tmpe.Next != nil {
		tmpo.Next = tmpo.Next.Next
		tmpo = tmpo.Next
		tmpe.Next = tmpe.Next.Next
		tmpe = tmpe.Next
	}
	tmpo.Next = even
	return odd
}
