package main

func main() {

}

func isOneBitCharacter(bits []int) bool {
	bitPos := 0
	for bitPos < len(bits)-1 {
		if bits[bitPos] == 1 {
			bitPos++
		}
		bitPos++
	}
	return bitPos == len(bits)-1
}
