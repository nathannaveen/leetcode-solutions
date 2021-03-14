package main

func main() {

}

func swapNodes(head *ListNode, k int) *ListNode {
	first := head
	second := head
	findLength := head
	lengthOfLinkedList := 0

	for findLength != nil {
		lengthOfLinkedList++
		findLength = findLength.Next
	}

	for i := 0; i < lengthOfLinkedList-k; i++ {
		second = second.Next
	}

	for i := 0; i < k-1; i++ {
		first = first.Next
	}

	second.Val, first.Val = first.Val, second.Val
	return head
}
