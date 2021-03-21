package main

func main() {

}

func killProcess(pid []int, ppid []int, kill int) []int {
	res := []int{kill}

	for i, i2 := range ppid {
		for _, re := range res {
			if i2 == re {
				res = append(res, pid[i])
				break
			}
		}
		numberCounter := 0
		for _, re := range res {
			if re == i2 {
				numberCounter++
			}
		}
		if numberCounter != 1 {
			res = res[:len(res)-1]
		}
	}
	return res
}
