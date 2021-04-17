package main

func main() {

}

func majorityElement(nums []int) []int {
	m := make(map[int]int)
	res := []int{}

	for _, num := range nums {
		m[num]++
		if m[num] == len(nums)/2+1 {
			res = append(res, num)
		}
	}
	return res
}
