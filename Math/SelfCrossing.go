package main

func main() {
}

func isSelfCrossing(x []int) bool {
	upDown := []int{}
	updownCounter := 0
	leftrightCounter := 0
	leftRight := []int{}
	side := 0
	leftRight = append(leftRight, 0)

	for _, i2 := range x {
		if side == 0 {
			if len(upDown) >= 2 && upDown[len(upDown)-2] <= updownCounter+i2 && len(leftRight) >= 3 &&
				leftRight[len(leftRight)-1] < leftRight[len(leftRight)-3] {
				return true
			}
			if updownCounter+i2 == 0 && len(upDown) >= 2 && leftRight[len(leftRight)-1] == leftRight[len(leftRight)-2] {
				return true
			}
			updownCounter += i2
			upDown = append(upDown, i2)
			side++
		} else if side == 1 {
			if len(leftRight) >= 3 && leftRight[len(leftRight)-3] >= leftrightCounter-i2 &&
				leftrightCounter > leftRight[len(leftRight)-3] && upDown[len(upDown)-1] >= upDown[len(upDown)-2] {
				return true
			}
			if len(leftRight) >= 3 && leftrightCounter-i2 <= leftRight[len(leftRight)-3] &&
				leftRight[len(leftRight)-1] > leftRight[len(leftRight)-3] {
				return true
			}
			leftrightCounter -= i2
			leftRight = append(leftRight, i2)
			side++
		} else if side == 2 {
			if len(upDown) >= 3 && upDown[len(upDown)-3] >= updownCounter-i2 {
				return true
			}
			upDown = append(upDown, i2)
			updownCounter -= i2
			side++
		} else {
			if len(leftRight) >= 2 && leftRight[len(leftRight)-2] <= leftrightCounter+i2 &&
				upDown[len(upDown)-1] <= upDown[len(upDown)-2] {
				return true
			}

			leftRight = append(leftRight, i2)
			leftrightCounter += i2
			side = 0
		}
	}

	return false
}
