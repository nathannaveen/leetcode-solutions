package main

import "math"

func main() {

}

func maxSatisfied(customers []int, grumpy []int, X int) int {
	max := 0
	res := 0
	temp := 0

	for i := 0; i < X; i++ {
		res, temp = AddToResOrTemp(customers, grumpy, i, res, temp)
	}

	max = temp

	for i := X; i < len(customers); i++ {
		if grumpy[i-X] == 1 {
			temp -= customers[i-X]
		}

		res, temp = AddToResOrTemp(customers, grumpy, i, res, temp)
		max = int(math.Max(float64(max), float64(temp)))
	}

	return res + max
}

func AddToResOrTemp(customers []int, grumpy []int, i int, res int, temp int) (int, int) {
	if grumpy[i] == 0 {
		res += customers[i]
	} else {
		temp += customers[i]
	}
	return res, temp
}
