package main

func main() {

}

func reorganizeString(S string) string {
	h := []string{}

	for _, i := range S {
		h = append(h, string(i))
	}

	left := 0
	right := len(h) - 1
	res := ""

	for left < right {
		res += h[left]
		left++
		if left < right {
			res += h[right]
		}
		right--
	}
	res += h[left]
	for i := 1; i < len(res); i++ {
		if string(res[i]) == string(res[i-1]) {
			return ""
		}
	}
	return res
}
