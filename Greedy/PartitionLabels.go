package main

func main() {

}

func partitionLabels(S string) []int {
	starts := make(map[int32]int)
	ends := make(map[int32]int)
	start, end := 0, 0
	res := []int{}

	for i, i2 := range S {
		if _, ok := starts[i2]; ok {
			ends[i2] = i
		} else {
			starts[i2] = i
		}
	}

	for i, i2 := range S {
		if ends[i2] > end {
			end = ends[i2]
		}
		if i == end {
			res = append(res, -1*(start-end)+1)
			start = i + 1
		}
	}
	return res
}
