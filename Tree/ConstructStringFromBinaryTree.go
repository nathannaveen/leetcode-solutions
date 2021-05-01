package main

import "strconv"

func main() {

}

func tree2str(t *TreeNode) string {
	stack := []*TreeNode{t}
	res := ""
	counter := 0

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != nil {
			if counter != 0 && res[len(res)-1] == ')' {
				res += "("
			}
			res += strconv.Itoa(pop.Val) + "("
			stack = append(stack, pop.Right, pop.Left)
		} else {
			if res[len(res)-1] == '(' {
				res = res[:len(res)-1]
			} else {
				res += ")"
			}
		}
		counter++
	}

	return res
}
