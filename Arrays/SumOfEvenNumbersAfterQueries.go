package main

func main() {

}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	total := 0
	for _, num := range A {
		if num%2 == 0 {
			total += num
		}
	}

	res := make([]int, 0, len(queries))
	for _, i := range queries {
		value := i[0]
		if A[i[1]]%2 == 0 {
			total -= A[i[1]]
		}
		if (A[i[1]]+value)%2 == 0 {
			total += A[i[1]] + value
		}
		A[i[1]] = A[i[1]] + value
		res = append(res, total)
	}

	return res
}
