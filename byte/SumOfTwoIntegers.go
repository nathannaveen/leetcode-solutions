package main

func main() {

}

func getSum(a int, b int) int {
	for i := 0; i < b; i++ {
		a = a & 1
	}

	return a
}
