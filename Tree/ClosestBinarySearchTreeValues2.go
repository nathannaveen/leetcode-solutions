package main

func main() {

}

func closestKValues(root *TreeNode, target float64, k int) []int {
	eachNumber := []int{}
	differences := []float64{}
	stack := []*TreeNode{root}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pop != nil {
			differences = append(differences, absFloat(float64(pop.Val)-target))
			eachNumber = append(eachNumber, pop.Val)
			stack = append(stack, pop.Left, pop.Right)
		}
	}

	for i := 1; i < len(differences); i++ {
		if i >= 1 && differences[i-1] > differences[i] {
			eachNumber[i], eachNumber[i-1] = eachNumber[i-1], eachNumber[i]
			differences[i], differences[i-1] = differences[i-1], differences[i]
			i -= 2
		}
	}

	return eachNumber[:k]
}

func absFloat(a float64) float64 {
	if a > 0 {
		return a
	}
	return -a
}
