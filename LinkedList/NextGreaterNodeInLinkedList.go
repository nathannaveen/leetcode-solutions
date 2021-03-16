package main

func main() {

}

func nextLargerNodes(head *ListNode) []int {
	res := []int{}

	for head != nil {
		l, max := head, 0

		for l.Next != nil {
			l = l.Next
			if l.Val > head.Val {
				max = l.Val
				break
			}
		}
		res = append(res, max)
		head = head.Next
	}
	return res
}
