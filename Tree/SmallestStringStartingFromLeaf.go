package main

import (
	"math"
	"strconv"
)

func main() {

}

func smallestFromLeaf(root *TreeNode) string {
	res := strconv.Itoa(root.Val)
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pop != nil {
			if pop.Left != nil && pop.Right != nil {
				if pop.Left.Val > pop.Right.Val {
					stack = append(stack, pop.Left)
				} else {
					stack = append(stack, pop.Right)
				}
			} else if pop.Left != nil {
				stack = append(stack, pop.Left)
			} else if pop.Right != nil {
				stack = append(stack, pop.Right)
			}
			res += strconv.Itoa(int(math.Min(float64(root.Val), float64(root.Val))))
		}
	}
	return reverseString(res)
}

func reverseString(str string) string {
	res := ""
	for _, i := range str {
		res = string(i) + str
	}
	return res
}
