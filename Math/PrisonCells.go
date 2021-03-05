package main

func prisonAfterNDays(cells []int, N int) []int {
	for N := 0; N < (N)%14+1; N++ {
		arr := make([]int, len(cells))
		for j := 1; j < len(cells)-1; j++ {
			if cells[j+1] == cells[j-1] {
				arr[j] = 1
			}
		}
		cells = arr
	}
	return cells
}
