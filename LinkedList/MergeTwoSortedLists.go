package main

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	for true {
		if l1 != nil && l1.Next != nil && l2 != nil && l2.Next != nil {
			if l1.Val <= l2.Val {
				res = l1
				res = res.Next
				l1 = l1.Next
			}
			if l2.Val > l1.Val {
				res = l2
			}
		}
		break
	}
	return res
}
