package main

func main() {
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	res := 0
	for i := range worker {
		for j := 0; j < len(difficulty); j++ {
			if j > 0 && difficulty[j] > worker[i] && difficulty[j-1] <= worker[i] {
				res += profit[j-1]
				break
			} else if j == 0 && difficulty[j] >= worker[i] {
				res += profit[0]
				break
			} else if j == len(difficulty)-1 && difficulty[j] <= worker[i] {
				res += profit[j]
			}
		}
	}
	return res
}
