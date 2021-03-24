package main

func main() {

}

func advantageCount(A []int, B []int) []int {
	res := []int{}
	for i, i2 := range B {
		minimum := 1000000001

		for _, i3 := range A {
			if i3 < minimum && i3 > i2 {
				minimum = i3
			}
		}
		for _, i3 := range A {
			if i3 == minimum {
				i3 = -1
			}
		}

		if minimum != 1000000001 {
			res = append(res, minimum)
		}
		B[i] = -1
	}
	//sort.Ints(B)
	//for _, i2 := range A {
	//	if i2 != -1 {
	//		res = append(res, i2)
	//	}
	//}

	return res
}
