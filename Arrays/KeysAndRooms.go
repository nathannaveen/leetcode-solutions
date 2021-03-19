package main

func canVisitAllRooms(rooms [][]int) bool {
	h := make([]bool, len(rooms))
	stack := []int{}
	for _, i := range rooms[0] {
		stack = append(stack, i)
	}
	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, i := range rooms[pop] {
			if i != pop && i != 0 {
				stack = append(stack, i)
			}
		}
		h[pop] = true
	}

	for i, b := range h {
		if i != 0 && b == false {
			return false
		}
	}
	return true
}
