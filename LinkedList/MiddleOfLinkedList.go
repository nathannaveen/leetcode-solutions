package main

import (
	"math"
)

func main() {

}

func middleNode(head *ListNode) *ListNode {
	cur := head
	counter := 0

	for cur != nil && cur.Next != nil {
		counter++
		cur = cur.Next
	}

	middle := int(math.Floor(float64(counter) / 2))

	if counter%2 == 0 {
		middle += 1
	}

	for i := 0; i < middle; i++ {
		head = head.Next
	}

	return head
}
