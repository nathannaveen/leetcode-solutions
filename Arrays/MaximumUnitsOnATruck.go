package main

func main() {

}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	res := 0
	for i := 1; i < len(boxTypes); i++ {
		if i >= 1 && boxTypes[i-1][0]*boxTypes[i-1][1] > boxTypes[i][0]*boxTypes[i][1] {
			boxTypes[i], boxTypes[i-1] = boxTypes[i-1], boxTypes[i]
		}
	}

	for _, boxType := range boxTypes {
		if boxType[0] < truckSize {
			res += boxType[0] * boxType[1]
			truckSize -= boxType[0]
		}
	}
	return res
}
