package main

func main() {

}

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	res := 0
	for _, num := range nums {
		sum += num
	}

	for sum != goal {
		if sum > goal {
			if sum-limit >= goal {
				sum -= limit
			} else {
				sum = goal
			}
		} else {
			if sum+limit <= goal {
				sum += limit
			} else {
				sum = goal
			}
		}
		res++
	}
	return res
}
