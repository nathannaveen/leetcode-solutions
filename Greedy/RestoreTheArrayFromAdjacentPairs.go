package main

type positionOfNumber struct {
	n         int
	positions []int
}

//func restoreArray(adjacentPairs [][]int) []int {
//	res := make([]int, len(adjacentPairs)+1)
//	arr := []positionOfNumber{}
//	m := make(map[int]int)
//	temp := 0
//	counter := 0
//	prev := -100001
//
//	for i, pair := range adjacentPairs {
//		g(m, pair, arr, i, 0)
//		g(m, pair, arr, i, 1)
//	}
//
//	for i, number := range arr {
//		if len(number.positions) == 1 {
//			temp = i
//			break
//		}
//	}
//
//	prev = arr[temp].positions[0]
//
//	for len(arr[temp].positions) != 1 || counter == 0 {
//
//		res = append(res, prev)
//		arr[temp].positions = append(arr[temp].positions[:0], arr[temp].positions[1:]...)
//		prev =
//
//		counter++
//	}
//
//	return res
//}
//
//func g(m map[int]int, pair []int, arr []positionOfNumber, i, pos int) {
//	if _, ok := m[pair[pos]]; ok {
//		arr[m[pair[pos]]].positions = append(arr[m[pair[pos]]].positions, i)
//	} else {
//		arr = append(arr, positionOfNumber{
//			n:         pair[pos],
//			positions: []int{i},
//		})
//		m[pair[pos]] = len(arr) - 1
//	}
//}
