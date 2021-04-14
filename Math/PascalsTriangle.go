package main

func main() {

}

func generate(numRows int) [][]int {
	res := [][]int{}

	for i := 0; i < numRows; i++ {
		h := make([]int, i+1)
		h[0], h[len(h)-1] = 1, 1
		for j := 1; j < len(h)-1; j++ {
			h[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, h)
	}

	return res
}
