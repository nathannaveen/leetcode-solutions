package main

func main() {

}

func splitListToParts(head *ListNode, k int) []*ListNode {
	res := []*ListNode{}
	length := 0
	cur := head

	for cur != nil {
		cur = cur.Next
		length++
	}
	cur = head

	g := length / k
	h := length % k

	for i := 0; i < k; i++ {
		l := &ListNode{}
		f := g

		if h > 0 {
			h--
			f++
		}
		if cur != nil {
			l = cur
			cur = cur.Next
		}
		for j := 1; j < f; j++ {
			l.Next = cur
			cur = cur.Next
		}

		if l.Val == 0 && l.Next == nil {
			res = append(res, nil)
		} else {
			res = append(res, l)
		}
	}

	return res
}
