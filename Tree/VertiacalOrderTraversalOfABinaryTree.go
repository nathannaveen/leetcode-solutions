package main

func main() {

}

func verticalTraversal(root *TreeNode) [][]int {
	h := [][]int{{root.Val, 0}}
	res := [][]int{}

	stack := []*TreeNode{root}
	stackX := []int{0}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		popX := stackX[len(stackX)-1]
		stack = stack[:len(stack)-1]
		stackX = stackX[:len(stackX)-1]

		if pop.Right != nil {
			stack = append(stack, pop.Right)
			stackX = append(stackX, popX+1)
			h = append(h, []int{pop.Right.Val, popX + 1})
		}
		if pop.Left != nil {
			stack = append(stack, pop.Left)
			stackX = append(stackX, popX-1)
			h = append(h, []int{pop.Left.Val, popX - 1})
		}
	}
	//  || (h[i-1][1] == h[i][1] && h[i-1][0] > h[i][0])
	for i := 1; i < len(h); i++ {
		if i >= 1 && (h[i-1][1] > h[i][1]) {
			h[i-1], h[i] = h[i], h[i-1]
			i -= 2
		}
	}
	arr := []int{h[0][0]}
	for i := 1; i < len(h); i++ {
		if h[i][1] == h[i-1][1] {
			arr = append(arr, h[i][0])
		} else {
			res = append(res, arr)
			arr = []int{h[i][0]}
		}
	}
	res = append(res, arr)
	return res
}
