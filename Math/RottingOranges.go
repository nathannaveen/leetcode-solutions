package main

func main() {

}

func orangesRotting(grid [][]int) int {
	minutes := 0
	rotting := true

	for rotting {
		rotting = false

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == 2 {
					//fmt.Println(i, j)
					if i > 0 && grid[i-1][j] == 1 {
						//fmt.Println("taco1")
						grid[i-1][j] = 3
						rotting = true
					}
					if i < len(grid)-1 && grid[i+1][j] == 1 {
						//fmt.Println("taco2")
						grid[i+1][j] = 3
						rotting = true
						//fmt.Println(grid[0])
						//fmt.Println(grid[1])
						//fmt.Println(grid[2])
					}
					if j > 0 && grid[i][j-1] == 1 {
						//fmt.Println("taco3")
						grid[i][j-1] = 3
						rotting = true
					}
					if j < len(grid)-1 && grid[i][j+1] == 1 {
						//fmt.Println("taco4")
						grid[i][j+1] = 3
						rotting = true
						//fmt.Println(grid[0])
						//fmt.Println(grid[1])
						//fmt.Println(grid[2])
					}
				}
			}
		}
		containsOne := false
		min := false
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == 3 {
					if !min {
						minutes++
						min = true
					}
					grid[i][j] = 2
				} else if grid[i][j] == 1 {
					containsOne = true
				}
			}
		}
		if !containsOne && minutes == 0 {
			return minutes
		}
		if !containsOne {
			return minutes + 1
		}

		if !rotting {
			return -1
		}

	}

	return minutes
}
