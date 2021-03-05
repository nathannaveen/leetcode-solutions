package main

import "math"

func main() {

}

func isPalindrome(head *ListNode) bool {
	j := []int{}
	counter := 0
	h := head
	for head != nil && head.Next != nil {
		counter++
		head = head.Next
	}

	middle := int(math.Floor(float64(counter) / 2))

	if counter%2 == 1 {
		middle += 1
	}

	for i := 0; i < middle; i++ {
		h = h.Next
	}

	for i := middle; i < counter; i++ {
		j = append(j, h.Val)
		h = h.Next
	}
	return false
}
