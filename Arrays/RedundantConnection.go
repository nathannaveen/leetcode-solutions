package main

func main() {

}

func findRedundantDirectedConnection(edges [][]int) []int {
	res := make([]int, 2)
	h := []int{}

	for _, edge := range edges {
		if containsIntArray(h, edge[0]) && containsIntArray(h, edge[1]) {
			res[0], res[1] = edge[0], edge[1]
			return res
		} else {
			if !containsIntArray(h, edge[0]) {
				h = append(h, edge[0])
			}
			if !containsIntArray(h, edge[1]) {
				h = append(h, edge[1])
			}
		}
	}

	return res
}

func containsIntArray(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
