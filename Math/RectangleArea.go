package main

func main() {

}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	ABCDRectangle := (D - B) * (C - A) // the first square
	EFGHRectangle := (H - F) * (G - E) // the second square
	sum := ABCDRectangle + EFGHRectangle

	left, right := max(A, E), min(G, C)
	up, down := min(D, H), max(F, B)

	if right > left && up > down {
		sum -= (right - left) * (up - down) // overlap
	}
	return sum
}
