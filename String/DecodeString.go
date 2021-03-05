package main

func main() {

}

func decodeString(s string) string {
	res := ""
	h := []int{}
	for _, i2 := range s {
		if i2 >= 97 && i2 <= 122 {
			if len(h) == 0 {
				res += string(i2)
			} else {

			}
		}
	}

	return ""
}
