package main

func main() {
	removeElements(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
		},
	}, 1)
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	//
	//if head != nil && head.Val == val && head.Next == nil {
	//	head = head.Next
	//}
	return head
}
