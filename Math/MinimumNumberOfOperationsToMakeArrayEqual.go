package main

func main() {

}

func minOperations(n int) int {
	if n%2 == 1 {
		return (n >> 1) * ((n >> 1) + 1)
	} else {
		return (n >> 1) * (n >> 1)
	}
}
