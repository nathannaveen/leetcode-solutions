package main

func main() {

}

func closestValue(root *TreeNode, target float64) int {
	stack := []*TreeNode{root}
	minimum := float64(1000000000)
	number := 0

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != nil {
			h := abs(target - float64(pop.Val))
			if h < minimum {
				number = pop.Val
				minimum = h
			}
			stack = append(stack, pop.Left, pop.Right)
		}
	}
	return number
}
