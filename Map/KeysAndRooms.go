package main

func main() {

}

func canVisitAllRooms(rooms [][]int) bool {
	h := []int{}
	g := make([]bool, len(rooms))
	g[0] = true

	for _, i := range rooms[0] {
		h = append(h, i)
	}

	for len(h) != 0 {
		x := 0
		x, h = h[len(h)-1], h[:len(h)-1]
		g[x] = true
		for _, i := range rooms[x] {
			h = append(h, i)
		}
	}

	for _, i := range g {
		if i == false {
			return false
		}
	}
	return true
}
