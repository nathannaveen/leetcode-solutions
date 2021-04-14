package main

func main() {

}

func getRow(rowIndex int) []int {
	res := []int{}

	for i := 0; i <= rowIndex; i++ {
		temp := make([]int, len(res))
		copy(temp, res)
		res = append(res, 1)
		for j := 1; j < len(res)-1; j++ {
			res[j] = temp[j-1] + temp[j]
		}
	}
	return res
}
