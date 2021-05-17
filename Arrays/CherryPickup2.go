package main

import "fmt"

func main() {

}

type point struct {
	x int
	y int
}

func cherryPickup(grid [][]int) int {
	arr := [][]int{make([]int, len(grid[0]))}
	arr[0][0] = grid[0][0]
	queue := []point{{y: 0, x: 0}}

	for i := 1; i < len(grid); i++ {
		arr = append(arr, make([]int, len(grid[0])))
		n := len(queue)

		for j := 0; j < n; j++ {
			pos := queue[0]
			queue = queue[1:]

			if pos.x-1 != -1 {
				funk(grid, arr, pos, queue, pos.x-1)
			}
			if pos.x+1 != len(arr[0]) {
				funk(grid, arr, pos, queue, pos.x+1)
			}

			funk(grid, arr, pos, queue, pos.x)
		}
	}

	queue = []point{{y: 0, x: len(grid[0])}}
	arr[0][len(grid[0])] = grid[0][len(grid[0])]

	for i := 1; i < len(grid); i++ {
		n := len(queue)

		for j := 0; j < n; j++ {
			pos := queue[0]
			queue = queue[1:]

			if pos.x-1 != -1 {
				funk2(grid, arr, pos, queue, pos.x-1)
			}
			if pos.x+1 != len(arr[0]) {
				funk2(grid, arr, pos, queue, pos.x+1)
			}

			funk2(grid, arr, pos, queue, pos.x)
		}
	}
	fmt.Println(arr)
	return 1
}

func funk(grid [][]int, arr [][]int, pos point, queue []point, posX int) {
	queue = append(queue, point{y: pos.y + 1, x: posX})
	arr[pos.y+1][posX] = max(grid[pos.y+1][posX]+arr[pos.y][posX], arr[pos.y+1][posX])
}

func funk2(grid [][]int, arr [][]int, pos point, queue []point, posX int) {
	queue = append(queue, point{y: pos.y + 1, x: posX})
	if arr[pos.y+1][posX] != 0 {
		arr[pos.y+1][posX] = max(arr[pos.y][posX], arr[pos.y+1][posX])
	} else {
		arr[pos.y+1][posX] = max(grid[pos.y+1][posX]+arr[pos.y][posX], arr[pos.y+1][posX])
	}
}
