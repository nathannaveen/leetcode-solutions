package main

func removeCoveredIntervals(intervals [][]int) int {
	counter := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][0] {
			g := intervals[i]

			for j := i - 1; j >= 0; j-- {
				if intervals[j][1] <= g[1] {
					intervals[j+1] = intervals[j]
					intervals[j] = g
					break
				}
				intervals[j+1] = intervals[j]
			}
		}
	}
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][0] <= intervals[i+1][0] {
			if intervals[i][1] >= intervals[i+1][1] || (intervals[i][1] <= intervals[i+1][1] &&
				intervals[i][0] == intervals[i+1][0]) {
				continue
			}
		}
		counter++
	}
	return counter + 1
}
