package main

func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {

	if x_center < x2 && x_center > x1 && y_center < y2 && y_center > y1 {
		// Checking whether the circle is in the middle of the square
		return true
	}

	if ((abs(x_center-x1)*abs(x_center-x1))+(abs(y_center-y1)*abs(y_center-y1)) <= radius*radius) ||
		((abs(x_center-x2)*abs(x_center-x2))+(abs(y_center-y2)*abs(y_center-y2)) <= radius*radius) ||
		((abs(x_center-x1)*abs(x_center-x1))+(abs(y_center-y2)*abs(y_center-y2)) <= radius*radius) ||
		((abs(x_center-x2)*abs(x_center-x2))+(abs(y_center-y1)*abs(y_center-y1)) <= radius*radius) {
		return true
	}

	if ((y1 <= y_center+radius && y1 >= y_center) || (y2 >= y_center-radius && y2 <= y_center)) &&
		x_center >= x1 && x_center <= x2 {
		return true
	}

	if y_center <= y2 && y_center >= y1 && ((x_center-radius <= x2 && x_center >= x2) ||
		(x_center+radius >= x1 && x_center <= x1)) {
		return true
	}

	return false
}

func abs(a int) int {
	if a > 0 {
		return a

	}
	return -a
}
