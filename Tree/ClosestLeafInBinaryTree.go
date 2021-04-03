package main

func main() {

}

func findClosestLeaf(root *TreeNode, k int) int {
	stack := []*TreeNode{root}
	stack2 := []int{0}
	hasPassedK := false
	counter := 0
	res := root.Val

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		pop2 := stack2[len(stack2)-1]
		stack = stack[:len(stack)-1]
		stack2 = stack2[:len(stack2)-1]
		if pop != nil {
			if hasPassedK {
				if pop.Right == nil && pop.Left == nil {
					counter = max(counter, pop2)
					res = pop.Val
				}
			} else if pop.Val == k {
				hasPassedK = true
			}
			stack2 = append(stack2, pop2+1, pop2+1)
			stack = append(stack, pop.Right, pop.Left)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
