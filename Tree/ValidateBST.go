package main

func main() {

}

func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != nil {
			if pop.Left != nil && pop.Left.Val >= pop.Val {
				return false
			} else if pop.Right != nil && pop.Right.Val <= pop.Val {
				return false
			}
			stack = append(stack, pop.Right, pop.Left)
		}
	}

	return true
}
