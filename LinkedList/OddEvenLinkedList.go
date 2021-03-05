package main

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

func main() {

}

func oddEvenList(head *ListNode) *ListNode {
	even := []int{}
	odd := []int{}
	g := head
	counter := 1
	for head != nil && head.Next != nil {
		if head.Val%2 == 1 {
			odd = append(odd, counter)
		} else {
			even = append(even, counter)
		}
		counter++
		head = head.Next
	}

	fmt.Println(odd, even)

	for _, i2 := range odd {
		fmt.Println(i2)
		g.Val = i2
		g = g.Next
	}
	for _, i2 := range even {
		fmt.Println(i2)
		g.Val = i2
		g = g.Next
	}
	return g
}
