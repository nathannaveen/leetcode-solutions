package main

func main() {

}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != nil {
			res = append(res, pop.Val)
			if pop.Right != nil {
				stack = append(stack, pop.Right)
			} else {
				stack = append(stack, pop.Left)
			}
		}
	}
	return res
}
