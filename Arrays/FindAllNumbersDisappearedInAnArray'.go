package main

func main() {

}

func findDisappearedNumbers(nums []int) []int {

	m := make(map[int]int)
	arr := []int{}

	for _, i := range nums {
		m[i]++
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			arr = append(arr, i)
		}
	}
	return arr[1:]
}
